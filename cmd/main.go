package main

import (
	"log"

	pkg "github.com/nicollasm/golang-postgre-importer-csv/pkg"
)

func main() {
	db, err := pkg.InitDB("dbconfig.json")
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

	pkg.ReadAndWriteToDB(db, "example_table", "example.csv")
}
