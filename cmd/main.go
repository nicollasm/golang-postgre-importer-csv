package main

import (
	"log"
	"sync"

	"github.com/nicollasm/golang-postgre-importer-csv/pkg"
	"github.com/progrium/macdriver"
	"github.com/progrium/macdriver/cocoa"
)

const maxInsertRetries = 3

func main() {
	go func() {
		macdriver.Main(func() {
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
			go pkg.ReadCSVIntoChannel(filePath, dataChan)

			var wg sync.WaitGroup
			for record := range dataChan {
				wg.Add(1)
				go pkg.InsertData(db, tableName, record, &wg, maxInsertRetries)
			}
			wg.Wait()
		})
	}()
}

func getUserInput(prompt string) string {
	resultChan := make(chan string)
	cocoa.T_Alert("AppName", prompt, "", "Ok", func(btn cocoa.NSButton) {
		resultChan <- btn.Title().String()
	})
	return <-resultChan
}
