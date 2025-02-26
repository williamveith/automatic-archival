package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

// Import reads all .bin files from importPath, unmarshals them as Year messages,
// extracts the line items, and inserts them into the 'invoices' table in batches.
func (cdb *Database) Import(importPath string) error {
	entries, err := os.ReadDir(importPath)
	if err != nil {
		return fmt.Errorf("failed to read directory %q: %w", importPath, err)
	}

	// Choose a chunk size (number of rows per insert).
	// Adjust this based on memory and performance needs.
	const chunkSize = 1000

	// We'll reuse a slice of rows for each chunk flush.
	var batchRows [][]interface{}

	// flushToDB calls cdb.InsertData to insert the current batch
	// and then resets the batch.
	flushToDB := func() error {
		if len(batchRows) == 0 {
			return nil
		}
		cdb.InsertData(batchRows)
		log.Printf("Flushed %d rows to DB.", len(batchRows))

		// Reset the slice for the next batch
		batchRows = batchRows[:0]
		return nil
	}

	totalInserted := 0 // Count total rows across all files

	for _, entry := range entries {
		// Only process regular files with .bin extension
		if !entry.Type().IsRegular() || filepath.Ext(entry.Name()) != ".bin" {
			continue
		}

		filePath := filepath.Join(importPath, entry.Name())

		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read file %q: %w", filePath, err)
		}

		// Unmarshal into a Year message (from schema.pb.go)
		var yearMsg Year
		if err := proto.Unmarshal(data, &yearMsg); err != nil {
			return fmt.Errorf("failed to unmarshal file %q as Year: %w", filePath, err)
		}

		fileRowCount := 0

		// For each Invoice in the Year, gather each LineItem
		for _, invoice := range yearMsg.Invoices {
			for _, li := range invoice.Lineitems {
				row := []interface{}{
					li.Date,        // date
					li.Time,        // time
					li.Prof,        // prof
					li.User,        // user
					li.Tool,        // tool
					li.Usage,       // usage
					li.Rate,        // rate
					li.Cost,        // cost
					li.AppliedCost, // applied_cost
				}
				batchRows = append(batchRows, row)
				fileRowCount++

				// If we hit chunkSize, flush to DB
				if len(batchRows) >= chunkSize {
					if err := flushToDB(); err != nil {
						return fmt.Errorf("failed flushing rows for %s: %w", filePath, err)
					}
				}
			}
		}

		totalInserted += fileRowCount
		log.Printf("File %q (year=%s) -> found %d line items", filePath, yearMsg.GetPeriod(), fileRowCount)
	}

	// After processing all files, flush any leftover rows
	if err := flushToDB(); err != nil {
		return fmt.Errorf("final flush error: %w", err)
	}

	log.Printf("Import completed. Total inserted rows: %d\n", totalInserted)
	return nil
}
