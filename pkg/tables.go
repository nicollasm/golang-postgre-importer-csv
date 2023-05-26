package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Cria a tabela se não existir
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

// Função para inserir dados na base de dados
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
		_, err := db.Exec(stmt, stringSliceToInterface(record)...)
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

// Converte uma fatia de strings em uma fatia de interfaces
func stringSliceToInterface(s []string) []interface{} {
	result := make([]interface{}, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}
