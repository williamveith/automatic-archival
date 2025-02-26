package database

import (
	"testing"
)

func TestImport(t *testing.T) {
	db := NewDatabase("/Users/main/Projects/Invoicing Archive/Invoices copy/db/invoices.sqlite3")
	db.Import("/Users/main/Projects/Invoicing Archive/Invoices copy/db/Archive")
}
