package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	sync "sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	dbName string
	schema string
	db     *sql.DB
	lock   sync.Mutex
}

const InsertSql string = `
INSERT INTO invoices (date, time, prof, user, tool, usage, rate, cost, applied_cost) 
VALUES %s
ON CONFLICT DO NOTHING
`

func NewDatabasePath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}

	projectRoot := filepath.Dir(filepath.Dir(exePath))
	dbPath := filepath.Join(projectRoot, "db", "invoices.sqlite3")

	return dbPath
}

func getSQLiteType(kind protoreflect.Kind) string {
	switch kind {
	case protoreflect.StringKind:
		return "TEXT NOT NULL"
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return "REAL NOT NULL"
	case protoreflect.Int32Kind, protoreflect.Int64Kind:
		return "INTEGER NOT NULL"
	default:
		return "TEXT NOT NULL" // Default fallback
	}
}

func generateSchemaFromProto(message proto.Message, tableName string) string {
	msgDescriptor := message.ProtoReflect().Descriptor()

	var schemaBuilder strings.Builder
	schemaBuilder.WriteString("PRAGMA journal_mode=WAL;\n")
	schemaBuilder.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", tableName))

	fieldCount := msgDescriptor.Fields().Len()
	columns := []string{} // Store column names for the UNIQUE constraint

	for i := 0; i < fieldCount; i++ {
		field := msgDescriptor.Fields().Get(i)
		columnName := string(field.Name())
		sqliteType := getSQLiteType(field.Kind())

		if i > 0 {
			schemaBuilder.WriteString(",\n")
		}
		schemaBuilder.WriteString(fmt.Sprintf("  %s %s", columnName, sqliteType))
		columns = append(columns, columnName)
	}

	if len(columns) > 0 {
		schemaBuilder.WriteString(",\n  UNIQUE(" + strings.Join(columns, ", ") + ")")
	}

	schemaBuilder.WriteString("\n);")
	return schemaBuilder.String()
}

func NewDatabase(dbPath string) *Database {
	var err error

	database := &Database{
		dbName: dbPath,
		schema: generateSchemaFromProto(&LineItem{}, "invoices"),
	}

	database.db, err = sql.Open("sqlite3", database.dbName)
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

		query := fmt.Sprintf(InsertSql, strings.Join(placeholders, ","))
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
