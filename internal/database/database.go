package database

import (
	"fmt"
	"log"
	"os"
	"strings"
	sync "sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	dbName string
	schema string
	db     *sqlx.DB
	lock   sync.Mutex
}

func readSqlStatement(filePath string) string {
	data, _ := os.ReadFile(filePath)
	return string(data)
}

func NewDatabase(dbPath string) *Database {
	var err error

	database := &Database{
		dbName: dbPath,
		schema: readSqlStatement("sql/schema.sql"),
	}

	database.db, err = sqlx.Open("sqlite3", database.dbName)
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}

	_, err = database.db.Exec(database.schema)
	if err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	return database
}

func (cdb *Database) InsertData(rows [][]interface{}) {
	cdb.lock.Lock()
	defer cdb.lock.Unlock()

	if len(rows) == 0 {
		fmt.Println("No data to insert.")
		return
	}

	tx, err := cdb.db.Begin()
	if err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}

	const batchSize = 1000
	insertSqlTemplate := readSqlStatement("sql/insert.sql")
	for i := 0; i < len(rows); i += batchSize {
		end := i + batchSize
		if end > len(rows) {
			end = len(rows)
		}

		batch := rows[i:end]

		// Build bulk insert query
		placeholders := []string{}
		values := []interface{}{}

		for _, row := range batch {
			ph := []string{}
			for range row {
				ph = append(ph, "?")
			}
			placeholders = append(placeholders, fmt.Sprintf("(%s)", strings.Join(ph, ",")))
			values = append(values, row...)
		}

		query := fmt.Sprintf(insertSqlTemplate, strings.Join(placeholders, ","))
		_, err := tx.Exec(query, values...)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Bulk insert failed: %v", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	fmt.Printf("Successfully inserted %d rows!\n", len(rows))
}
