package pkg

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	tableNameSchema = `CREATE TABLE IF NOT EXISTS %s (
		TELEFONE varchar(255),
		DATA_ATIVACAO varchar(255),
		DATA_RETIRADA varchar(255),
		STATUS_PRODUTO varchar(255),
		TECNOLOGIA varchar(255),
		TECNOLOGIA_BANDA varchar(255),
		DOCUMENTO varchar(255),
		TIPO_DOCUMENTO varchar(255),
		NOME_CLIENTE varchar(255),
		TIPO_CLIENTE varchar(255),
		TIPO_ENDERECO varchar(255),
		ENDERECO varchar(255),
		NUMERO_ENDERECO varchar(255),
		COMPL_ENDERECO varchar(255),
		CEP varchar(255),
		CEL_CONTATO varchar(255),
		FIXO_CONTATO varchar(255),
		EMAIL varchar(255),
		UF varchar(255),
		REGIAO varchar(255),
		BAIRRO varchar(255)
	);`
	insertSchema = `INSERT INTO %s (
			TELEFONE,
			DATA_ATIVACAO,
			DATA_RETIRADA,
			STATUS_PRODUTO,
			TECNOLOGIA,
			TECNOLOGIA_BANDA,
			DOCUMENTO,
			TIPO_DOCUMENTO,
			NOME_CLIENTE,
			TIPO_CLIENTE,
			TIPO_ENDERECO,
			ENDERECO,
			NUMERO_ENDERECO,
			COMPL_ENDERECO,
			CEP,
			CEL_CONTATO,
			FIXO_CONTATO,
			EMAIL,
			UF,
			REGIAO,
			BAIRRO
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
)

const batchSize = 5000

func CreateTable(db *sql.DB, tableName string) error {
	stmt := fmt.Sprintf(tableNameSchema, tableName)
	_, err := db.Exec(stmt)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Tabela criada com sucesso.")
	return nil
}

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
	count := 0
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
		count++
		if count >= batchSize {
			InsertData(db, tableName, data)
			data = nil
			count = 0
		}
	}
	if count > 0 {
		InsertData(db, tableName, data)
	}
}

func InsertData(db *sql.DB, tableName string, records [][]string) error {
	sqlStatement := fmt.Sprintf(insertSchema, tableName)

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for _, record := range records {
		if len(record) != 21 {
			log.Printf("Número incorreto de campos: %d. Esperado: 21. Pulando a linha", len(record))
			continue
		}
		_, err = stmt.Exec(record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9], record[10], record[11], record[12], record[13], record[14], record[15], record[16], record[17], record[18], record[19], record[20])
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 2)
			continue
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
