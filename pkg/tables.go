package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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
	);`, tableName)
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
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, tableName)

	for i := 0; i <= retries; i++ {
		_, err := db.Exec(stmt, record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9], record[10], record[11], record[12], record[13], record[14], record[15], record[16], record[17], record[18], record[19], record[20])
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 2)
			continue
		}
		break
	}
	return nil
}
