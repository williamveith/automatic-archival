package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/golang/snappy"

	"google.golang.org/protobuf/proto"
)

func parseFloatCommaSafe(s string) (float64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0.0, nil
	}
	// remove commas
	s = strings.ReplaceAll(s, ",", "")
	// parse as float64
	return strconv.ParseFloat(s, 64)
}

func (cdb *Database) Export(exportPath string) error {
	// Query all rows (adjust columns as needed).
	// If your table name is something else or your columns differ, change accordingly.
	rows, err := cdb.db.Query(`SELECT * FROM invoices`)

	if err != nil {
		log.Fatalf("failed to query rows: %v", err)
	}
	defer rows.Close()

	// Maps for grouping:
	//  year -> group -> *Invoice
	yearInvoices := make(map[string]map[string]*Invoice)

	for rows.Next() {
		var (
			dateStr, timeStr, prof, user, tool     string
			usageStr, rateStr, costStr, appliedStr string
		)
		err := rows.Scan(&dateStr, &timeStr, &prof, &user, &tool, &usageStr, &rateStr, &costStr, &appliedStr)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		// Now clean and parse the numeric fields:
		usage, err := parseFloatCommaSafe(usageStr)
		if err != nil {
			log.Fatalf("parse usage error: %v", err)
		}
		rate, err := parseFloatCommaSafe(rateStr)
		if err != nil {
			log.Fatalf("parse rate error: %v", err)
		}
		cost, err := parseFloatCommaSafe(costStr)
		if err != nil {
			log.Fatalf("parse cost error: %v", err)
		}
		appliedCost, err := parseFloatCommaSafe(appliedStr)
		if err != nil {
			log.Fatalf("parse applied_cost error: %v", err)
		}
		// Parse out the year from dateStr. For example,
		// if dateStr is "2021-09-01", parse with time.Parse.
		// Adjust layout if your date format differs.
		var year string
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			// fallback: maybe the date is in another format or missing
			// if you can't parse it, handle error or skip
			log.Printf("failed to parse date: %s, err: %v", dateStr, err)
			continue
		}
		year = fmt.Sprintf("%04d", t.Year())

		// Create a new LineItem
		lineItem := &LineItem{
			Date:        dateStr,
			Time:        timeStr,
			Prof:        prof,
			User:        user,
			Tool:        tool,
			Usage:       usage,
			Rate:        rate,
			Cost:        cost,
			AppliedCost: appliedCost,
		}

		// Initialize yearInvoices[year] map if needed
		if _, ok := yearInvoices[year]; !ok {
			yearInvoices[year] = make(map[string]*Invoice)
		}

		// The "group" in your proto is stored in the "prof" column, or you might store
		// it differently. Adjust as necessary:
		groupKey := prof

		// If there's no existing invoice for this group in this year, create one
		invoiceMap := yearInvoices[year]
		if _, ok := invoiceMap[groupKey]; !ok {
			invoiceMap[groupKey] = &Invoice{
				Period: year, // or something like "2021-09" if you want month-based
				Group:  groupKey,
				// usage, cost, etc. will be aggregated as we go
			}
		}

		inv := invoiceMap[groupKey]

		// Add the line item
		inv.Lineitems = append(inv.Lineitems, lineItem)

		// Aggregate usage and costs for the invoice
		inv.Usage += usage
		inv.Cost += cost
		inv.AppliedCost += appliedCost
	}
	// Check for row iteration error
	if err := rows.Err(); err != nil {
		log.Fatalf("rows iteration error: %v", err)
	}

	// Now build the top-level Year messages from the yearInvoices map
	// year -> []*pb.Invoice
	// Then compute usage/cost across all invoices in each year
	for y, invoiceMap := range yearInvoices {
		// Create a Year message
		yearMsg := &Year{
			Period: y,
		}

		// Move map -> slice
		for _, inv := range invoiceMap {
			// Summation at the year level
			yearMsg.Usage += inv.Usage
			yearMsg.Cost += inv.Cost
			yearMsg.AppliedCost += inv.AppliedCost
			// If you have a known or dynamic surcharge/tax, set them here; e.g.:
			// inv.Surcharge = computeSurcharge(inv.Cost)
			// inv.Tax       = computeTax(inv.Cost)
			// yearMsg.Surcharge += inv.Surcharge
			// yearMsg.Tax       += inv.Tax

			yearMsg.Invoices = append(yearMsg.Invoices, inv)
		}

		// Optionally compute or set surcharge/tax at the year level
		// yearMsg.Surcharge = ...
		// yearMsg.Tax       = ...

		// Marshal to binary
		outFile := filepath.Join(exportPath, fmt.Sprintf("%s.bin", y)) // e.g. "2021.bin"
		data, err := proto.Marshal(yearMsg)
		if err != nil {
			log.Fatalf("failed to marshal Year for %s: %v", y, err)
		}

		if err := saveSnappyProto(data, outFile); err != nil {
			log.Fatalf("failed to write Snappy compressed file %s: %v", outFile, err)
		}

		fmt.Printf("Wrote year file: %s\n", outFile)
	}

	fmt.Println("All done.")
	return nil
}

func saveSnappyProto(data []byte, outFile string) error {
	// Create file with .snappy extension
	outFile += ".snappy"
	file, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("failed to create compressed file: %w", err)
	}
	defer file.Close()

	// Compress data using Snappy
	compressedData := snappy.Encode(nil, data)

	// Write compressed data to file
	_, err = file.Write(compressedData)
	if err != nil {
		return fmt.Errorf("failed to write snappy compressed data: %w", err)
	}

	fmt.Printf("Snappy compressed file written: %s\n", outFile)
	return nil
}
