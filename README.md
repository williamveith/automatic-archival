# Invoices Import/Export Tool

## Purpose and Motivation

As invoice data accumulates over time, storing it in a traditional relational database can become **unwieldy**—especially if you need to **share**, **archive**, or **query** data flexibly. Spreadsheets like **Excel** quickly become **infeasible** due to row limits, and **databases are not always available** or easy to maintain.

This tool solves that problem by:

✅ **Serializing invoice data using Protocol Buffers (Protobuf)** for long-term, compact archival storage.  
✅ **Ensuring that each year’s data is self-contained and complete** by bundling invoices into a single `.bin` file per year.  
✅ **Halving storage size** compared to databases, reducing thousands of small files into **one file per year**.  
✅ **Providing instant, ultra-available re-import** by converting `.bin` files back into a SQLite database on demand.  
✅ **Eliminating database management overhead**, since data can be stored in `.bin` archives and reconstructed when needed.

Instead of maintaining **400+ files per year**, this tool **reduces** them to **1 per year**. These `.bin` files are highly portable, compact, and **independent of any database system**—ensuring long-term data durability.

---

## How It Works

1. **Export**: Converts an SQLite database into a **set of yearly `.bin` files** for efficient, long-term storage.
2. **Import**: Reconstructs a **fully functional SQLite database** from archived `.bin` files whenever needed.

This ensures that **each year is guaranteed to be complete** and that invoices remain accessible while being stored in a compact, serialized format.

---

## Table of Contents

- [Project Structure](#project-structure)
- [Requirements](#requirements)
- [Build and Run](#build-and-run)
  - [Building](#building)
  - [Running](#running)
- [Usage](#usage)
  - [Exporting Data to Bin Files](#exporting-data-to-bin-files)
  - [Importing Data from Bin Files](#importing-data-from-bin-files)
- [Development](#development)
  - [Protobuf Compilation](#protobuf-compilation)
  - [Cleaning Up](#cleaning-up)
  - [Code Archive](#code-archive)
- [License](#license)

---

## Project Structure

```sh
.
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   └── invoices
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   └── database
│       ├── database.go
│       ├── export.go
│       ├── export_test.go
│       ├── import.go
│       ├── import_test.go
│       ├── schema.pb.go
│       └── sql
│           ├── insert.sql
│           └── schema.sql
└── protobuf
    └── schema.proto
```

---

## Requirements

- **Go 1.23.2+**
- **protoc** (Protocol Buffers compiler) if you need to re-generate `schema.pb.go`.  
- A Unix-like shell for running the `Makefile` commands (macOS, Linux, Git Bash on Windows, etc.).

---

## Build and Run

### Building

To compile the `invoices` binary:

```bash
make build
```

This will:

1. Create the `bin/` directory (if it doesn’t exist).  
2. Compile and place the binary (`invoices`) in `bin/`.  
3. Store the build logs in `bin/logs/build.log`.

### Running

After building, run:

```bash
make run
```

The tool will prompt you to **import** (`i`) or **export** (`e`).

---

## Usage

When running the application (`make run`), the tool asks:

1. **Would you like to import (i) or export (e)?**  
   - Type **`i`** to import data from `.bin` files into a SQLite database.  
   - Type **`e`** to export data from the SQLite database to `.bin` files.

### Exporting Data to Bin Files

1. Choose **`e`** at the prompt.  
2. **Enter the path** to your `invoices.sqlite3` database file (e.g. `db/invoices.sqlite3`).  
3. The tool creates an `Archive` folder (if missing) alongside the provided DB file path.  
4. It writes each year’s data to a **separate `.bin` file** (e.g. `Archive/2023.bin`).

You should see logs like:

```sh
Wrote year file: db/Archive/2023.bin
Wrote year file: db/Archive/2024.bin
...
Export finished successfully.
```

### Importing Data from Bin Files

1. Choose **`i`** at the prompt.  
2. **Enter the path** to the directory containing all your `.bin` files (e.g. `db/archive`).  
3. The tool **infers** the **parent** directory for the SQLite DB (e.g., `db/invoices.sqlite3`).  
4. It chunk‐loads the `.bin` files into your database, printing progress as it goes.

Sample output might look like:

```sh
File "db/archive/2023.bin" (year=2023) -> found 2500 line items
Flushed 1000 rows to DB.
Flushed 1000 rows to DB.
Flushed 500 rows to DB.
Import completed. Total inserted rows: 2500
Import finished successfully.
```

---

## Development

### Protobuf Compilation

If you modify `protobuf/schema.proto`, regenerate the Go code:

```bash
make proto
```

This runs:

```bash
protoc --go_out=. protobuf/schema.proto
```

which generates/overwrites `internal/database/schema.pb.go`.

### Cleaning Up

To remove the `bin/` folder (and your built binary/logs):

```bash
make clean
```

**Be careful**: If you have anything personal stored in `bin/`, this will remove it.

### Code Archive

If you want to create a `.tar.gz` file of **only** your tracked source code (ignoring `db/` and `bin/`), run:

```bash
make code-export
```

This creates `invoice-archive-import-export.tar.gz` with your project code, **excluding** local data.

---

## Why This Works

- **Efficient**: Uses Protobuf for space-efficient serialization.  
- **Database-Free Storage**: Data is stored **outside** of any live database.  
- **Portable**: `.bin` files can be stored, transferred, or backed up easily.  
- **Ultra-Available**: Need an invoice report? Just re-import and query.  
- **Compression-Friendly**: `.bin` files are naturally compact, but can be compressed further if needed.

---

## License

[Insert your license information here.]

---

With this README, anyone can:

1. Clone/download your repository.  
2. Run `make build && make run`.  
3. Interactively **import** or **export** invoice data to/from `.bin` files.  

🚀 **This ensures invoices remain highly available, durable, and easy to work with—without database headaches!**
