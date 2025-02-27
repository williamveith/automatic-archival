package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang/snappy"
	"google.golang.org/protobuf/proto"
)

func (cdb *Database) Import(importPath string) error {
	entries, err := os.ReadDir(importPath)
	if err != nil {
		return fmt.Errorf("failed to read directory %q: %w", importPath, err)
	}

	const chunkSize = 1000

	var batchRows [][]interface{}

	flushToDB := func() error {
		if len(batchRows) == 0 {
			return nil
		}
		cdb.InsertData(batchRows)
		log.Printf("Flushed %d rows to DB.", len(batchRows))

		batchRows = batchRows[:0]
		return nil
	}

	totalInserted := 0

	for _, entry := range entries {
		if !entry.Type().IsRegular() || filepath.Ext(entry.Name()) != ".snappy" {
			continue
		}

		filePath := filepath.Join(importPath, entry.Name())

		compressedData, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read file %q: %w", filePath, err)
		}

		data, err := snappy.Decode(nil, compressedData)
		if err != nil {
			return fmt.Errorf("failed to decompress file %q: %w", filePath, err)
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
					li.Date,    // date
					li.Time,    // time
					li.Prof,    // prof
					li.User,    // user
					li.Tool,    // tool
					li.Usage,   // usage
					li.Rate,    // rate
					li.Cost,    // cost
					li.Applied, // applied
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
