package pkg

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// Cria a tabela se ela não existir
func CreateTable(db *sql.DB, tableName string) error {
	stmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
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
	)`, tableName)
	_, err := db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Tabela criada com sucesso.")
	return nil
}

// Lê o CSV e insere os dados na tabela
func InsertDataFromCSV(db *sql.DB, tableName, filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		insertData(db, tableName, record)
	}
	fmt.Println("Dados inseridos com sucesso.")
	return nil
}

// Insere uma linha de dados na tabela
func insertData(db *sql.DB, tableName string, record []string) error {
	stmt := fmt.Sprintf(`INSERT INTO %s(
		TELEFONE, DATA_ATIVACAO, DATA_RETIRADA, STATUS_PRODUTO, 
		TECNOLOGIA, TECNOLOGIA_BANDA, DOCUMENTO, TIPO_DOCUMENTO, 
		NOME_CLIENTE, TIPO_CLIENTE, TIPO_ENDERECO, ENDERECO, 
		NUMERO_ENDERECO, COMPL_ENDERECO, CEP, CEL_CONTATO, 
		FIXO_CONTATO, EMAIL, UF, REGIAO, BAIRRO) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, tableName)
	_, err := db.Exec(stmt, record[0], record[1], record[2], record[3], record[4], record[5], record[6],
		record[7], record[8], record[9], record[10], record[11], record[12], record[13], record[14],
		record[15], record[16], record[17], record[18], record[19], record[20])
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}