package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <input.db> <output.json>\n", os.Args[0])
	}

	dbFile := os.Args[1]
	jsonFile := os.Args[2]

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM cards")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)

	for rows.Next() {
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}

		tableData = append(tableData, entry)
	}

	jsonData, err := json.MarshalIndent(tableData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(jsonFile, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data has been written to %s\n", jsonFile)
}
