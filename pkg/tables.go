package pkg

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func CreateTable(db *sql.DB, tableName string) error {
	stmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		//...
	)`, tableName)
	_, err := db.Exec(stmt)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Tabela criada com sucesso.")
	return nil
}

func InsertData(db *sql.DB, tableName string, record []string, wg *sync.WaitGroup, retries int) error {
	defer wg.Done()

	stmt := fmt.Sprintf(`INSERT INTO %s (
		//...
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, tableName)

	for i := 0; i <= retries; i++ {
		_, err := db.Exec(stmt, record...)
		if err != nil {
			if i == retries {
				log.Printf("Falhou após %d tentativas: %v\n", i+1, err)
				return err
			}
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		break
	}
	return nil
}

func ReadCSVIntoChannel(filepath string, dataChan chan []string) error {
	f, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			close(dataChan)
			break
		}
		if err != nil {
			log.Println(err)
			return err
		}
		if len(record) < 21 {
			log.Println("Registro com menos de 21 campos, ignorando...")
			continue
		}
		dataChan <- record
	}
	return nil
}
