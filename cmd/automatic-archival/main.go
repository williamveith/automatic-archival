package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/williamveith/automatic-archival/internal/database"
)

func importFromBin(dbPath string, binDir string) error {
	db := database.NewDatabase(dbPath)

	if err := db.Import(binDir); err != nil {
		return fmt.Errorf("import failed: %w", err)
	}
	return nil
}

func exportToBin(dbPath string, binDir string) error {
	db := database.NewDatabase(dbPath)

	if err := db.Export(binDir); err != nil {
		return fmt.Errorf("export failed: %w", err)
	}
	return nil
}

func cleanInput(input string) string {
	inputCleaned := strings.TrimSpace(input)
	return strings.Trim(inputCleaned, `"'`)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Would you like to import (i) or export (e)? ")
	choice, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	choice = strings.TrimSpace(strings.ToLower(choice))

	switch choice {
	case "i":
		fmt.Print("Enter the path to the directory containing the .bin files: ")
		binDir, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read bin directory path: %v", err)
		}
		binDir = cleanInput(binDir)

		parentDir := filepath.Dir(binDir)
		dbPath := filepath.Join(parentDir, "invoices.sqlite3")

		if err := importFromBin(dbPath, binDir); err != nil {
			log.Fatalf("Import failed: %v", err)
		}
		fmt.Println("Import finished successfully.")

	case "e":
		fmt.Print("Enter the path to the database (e.g., '/path/to/invoices.sqlite3'): ")
		dbPath, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read database path: %v", err)
		}
		dbPath = cleanInput(dbPath)

		dbDir := filepath.Dir(dbPath)
		archiveDir := filepath.Join(dbDir, "Archive")

		if err := os.MkdirAll(archiveDir, 0755); err != nil {
			log.Fatalf("Failed to create or verify Archive directory: %v", err)
		}

		if err := exportToBin(dbPath, archiveDir); err != nil {
			log.Fatalf("Export failed: %v", err)
		}
		fmt.Println("Export finished successfully.")

	default:
		fmt.Println("No valid choice made. Exiting.")
	}
}
