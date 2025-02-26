package database

import (
	"testing"
)

func TestExport(t *testing.T) {
	db := NewDatabase("/Users/main/Projects/Invoicing Archive/Invoices copy/db/invoices.sqlite3")
	db.Export("/Users/main/Projects/Invoicing Archive/Invoices copy/internal/database")
}
