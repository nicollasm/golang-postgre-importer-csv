package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nicollasm/golang-postgre-importer-csv/pkg"
)

const maxInsertRetries = 3

// Função para obter a entrada do usuário
func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}

// Lê um arquivo CSV e envia suas linhas para um canal
func ReadCSVIntoChannel(filepath string, dataChan chan []string) error {
	f, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	lineNum := 0
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
		if len(record) >= 21 {
			dataChan <- record
		} else {
			log.Printf("Registro com menos de 21 campos (linha %d), ignorando...", lineNum)
		}
		lineNum++
	}
	return nil
}

func main() {
	config := pkg.DBConfig{}

	config.Host = getUserInput("Insira o host do banco de dados:")
	config.Port = getUserInput("Insira a porta do banco de dados:")
	config.User = getUserInput("Insira o usuário do banco de dados:")
	config.Password = getUserInput("Insira a senha do banco de dados:")
	config.Database = getUserInput("Insira o nome do banco de dados:")
	tableName := getUserInput("Insira o nome da tabela:")
	filePath := getUserInput("Insira o caminho do arquivo CSV:")

	db, err := pkg.InitDB(config)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	err = pkg.CreateTable(db, tableName)
	if err != nil {
		log.Println(err)
		return
	}

	dataChan := make(chan []string)
	go ReadCSVIntoChannel(filePath, dataChan)

	var wg sync.WaitGroup
	for record := range dataChan {
		wg.Add(1)
		go pkg.InsertData(db, tableName, record, &wg, maxInsertRetries)
	}
	wg.Wait()
}
