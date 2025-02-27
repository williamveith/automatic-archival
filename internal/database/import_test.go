package database

import (
	"testing"
)

func TestImport(t *testing.T) {
	db := NewDatabase("db/invoices.sqlite3")
	db.Import("db/Archive")
}
