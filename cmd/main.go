package main

import (
	"log"

	pkg "github.com/nicollasm/golang-postgre-importer-csv/pkg"
)

func main() {
	db, err := pkg.InitDB(pkg.DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "nickas12",
		Database: "maindatabase",
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer db.Close()

	err = pkg.CreateTable(db, "example_table")
	if err != nil {
		log.Fatalln(err)
		return
	}

	pkg.ReadAndWriteToDB(db, "example_table", "example.csv") // ignorando o cabeçalho do CSV
}
