package main

import (
	"log"

	"github.com/nicollasm/golang-postgre-importer-csv/pkg"
	"github.com/progrium/macdriver"
	"github.com/progrium/macdriver/cocoa"
)

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

			err = pkg.InsertDataFromCSV(db, tableName, filePath)
			if err != nil {
				log.Println(err)
				return
			}
		})
	}()
}

func getUserInput(prompt string) string {
	resultChan := make(chan string)
	cocoa.T_Alert("AppName", prompt, "", "Ok", func(btn cocoa.NSButton) {
		resultChan <- btn.Title()
	})
	return <-resultChan
}
