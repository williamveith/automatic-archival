package database

import (
	"testing"
)

func TestExport(t *testing.T) {
	db := NewDatabase("db/invoices.sqlite3")
	db.Export("db/Archive2")
}
