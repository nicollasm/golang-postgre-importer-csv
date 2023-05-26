package pkg

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
	"sync"
)

func ReadCSV(csvName string) ([][]string, error) {
	f, err := os.Open(csvName)
	if err != nil {
		log.Println(err)
		return nil, err
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
			return nil, err
		}
		data = append(data, record)
	}
	return data, nil
}

func WriteToDB(db *sql.DB, tableName string, data [][]string) {
	var wg sync.WaitGroup
	for _, record := range data {
		wg.Add(1)
		go InsertData(db, tableName, record, &wg, 3)
	}
	wg.Wait()
}
