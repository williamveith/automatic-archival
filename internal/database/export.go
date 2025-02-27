package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang/snappy"

	"google.golang.org/protobuf/proto"
)

func (cdb *Database) Export(exportPath string) error {
	rows, err := cdb.db.Queryx(`SELECT * FROM invoices`)

	if err != nil {
		log.Fatalf("failed to query rows: %v", err)
	}
	defer rows.Close()
	yearInvoices := make(map[string]map[string]*Invoice)

	lineItem := &LineItem{}
	for rows.Next() {
		err := rows.StructScan(lineItem)
		if err != nil {
			log.Printf("Failed to StructScan: %v", err)
		}

		var year string
		t, err := time.Parse("2006-01-02", lineItem.Date)
		if err != nil {
			log.Printf("failed to parse date: %s, err: %v", lineItem.Date, err)
			continue
		}
		year = fmt.Sprintf("%04d", t.Year())

		if _, ok := yearInvoices[year]; !ok {
			yearInvoices[year] = make(map[string]*Invoice)
		}

		invoiceMap := yearInvoices[year]
		if _, ok := invoiceMap[lineItem.Prof]; !ok {
			invoiceMap[lineItem.Prof] = &Invoice{
				Period: year,
				Group:  lineItem.Prof,
			}
		}

		inv := invoiceMap[lineItem.Prof]
		inv.Lineitems = append(inv.Lineitems, lineItem)

		inv.Usage += lineItem.Usage
		inv.Cost += lineItem.Cost
		inv.Applied += lineItem.Applied
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("rows iteration error: %v", err)
	}

	for y, invoiceMap := range yearInvoices {
		yearMsg := &Year{
			Period: y,
		}

		for _, inv := range invoiceMap {
			yearMsg.Usage += inv.Usage
			yearMsg.Cost += inv.Cost
			yearMsg.Applied += inv.Applied
			yearMsg.Surcharge += inv.Surcharge
			yearMsg.Tax += inv.Tax
			yearMsg.Invoices = append(yearMsg.Invoices, inv)
		}

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
	outFile += ".snappy"
	file, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("failed to create compressed file: %w", err)
	}
	defer file.Close()

	compressedData := snappy.Encode(nil, data)
	_, err = file.Write(compressedData)
	if err != nil {
		return fmt.Errorf("failed to write snappy compressed data: %w", err)
	}

	fmt.Printf("Snappy compressed file written: %s\n", outFile)
	return nil
}
