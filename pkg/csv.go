package pkg

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ReadAndWriteToDB(db *sql.DB, tableName, csvName string) {
	f, err := os.Open(csvName)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.LazyQuotes = true

	var data [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return
		}
		data = append(data, record)
		InsertData(db, tableName, data)
		data = nil
	}
}
