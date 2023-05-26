package pkg

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	batchSize       = 1000
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

func ReadAndWriteToDB(db *sql.DB, fileName string, tableName string) {
	err := createTable(db, tableName)
	if err != nil {
		log.Println(err) // Alterado de log.Fatal(err)
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err) // Alterado de log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.LazyQuotes = true

	tx, err := db.Begin()
	if err != nil {
		log.Println(err) // Alterado de log.Fatal(err)
	}

	stmt, err := tx.Prepare(fmt.Sprintf(insertSchema, tableName))
	if err != nil {
		log.Println(err) // Alterado de log.Fatal(err)
	}

	linesBatch := make([][]string, 0, batchSize)
	lineCount := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err) // Alterado de log.Fatal(err)
			continue
		}

		// Verifique se a linha tem a quantidade correta de campos
		if len(record) != 21 {
			log.Printf("A linha %d tem um número incorreto de campos e será ignorada\n", lineCount+1)
			continue
		}

		linesBatch = append(linesBatch, record)

		// Se já temos bastante linhas, inserimos no banco de dados
		if len(linesBatch) == batchSize {
			insertBatch(linesBatch, stmt)
			linesBatch = linesBatch[:0] // Limpe o lote para o próximo
		}

		lineCount++
	}

	// Insira quaisquer linhas restantes que não foram inseridas
	if len(linesBatch) > 0 {
		insertBatch(linesBatch, stmt)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err) // Alterado de log.Fatal(err)
	}
}

// Função para inserir um lote de linhas
func insertBatch(linesBatch [][]string, stmt *sql.Stmt) {
	for _, line := range linesBatch {
		_, err := stmt.Exec(stringSliceToInterfaceSlice(line)...)
		if err != nil {
			log.Printf("Erro ao inserir a linha no banco de dados: %v\n", err)
		}
	}
}

func stringSliceToInterfaceSlice(stringSlice []string) []interface{} {
	interfaceSlice := make([]interface{}, len(stringSlice))
	for i, d := range stringSlice {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}

func createTable(db *sql.DB, tableName string) error {
	stmt := fmt.Sprintf(tableNameSchema, tableName)
	_, err := db.Exec(stmt)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
