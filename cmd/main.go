package main

import (
	"bufio"
	"fmt"
	"github.com/nicollasm/golang-postgre-importer-csv/pkg"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Insira o host do banco de dados:")
	host, _ := reader.ReadString('\n')

	fmt.Println("Insira a porta do banco de dados:")
	port, _ := reader.ReadString('\n')

	fmt.Println("Insira o usuário do banco de dados:")
	user, _ := reader.ReadString('\n')

	fmt.Println("Insira a senha do banco de dados:")
	password, _ := reader.ReadString('\n')

	fmt.Println("Insira o nome do banco de dados:")
	database, _ := reader.ReadString('\n')

	fmt.Println("Insira o nome da tabela:")
	tableName, _ := reader.ReadString('\n')

	fmt.Println("Insira o caminho do arquivo CSV:")
	filePath, _ := reader.ReadString('\n')

	cfg := pkg.DBConfig{
		Host:     strings.TrimSpace(host),
		Port:     strings.TrimSpace(port),
		User:     strings.TrimSpace(user),
		Password: strings.TrimSpace(password),
		Database: strings.TrimSpace(database),
	}

	db, err := pkg.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = pkg.CreateTable(db, strings.TrimSpace(tableName))
	if err != nil {
		log.Fatal(err)
	}

	err = pkg.InsertDataFromCSV(db, strings.TrimSpace(tableName), strings.TrimSpace(filePath))
	if err != nil {
		log.Fatal(err)
	}
}
