package main

import (
	"fmt"
	"log"
	"time"

	pkg "github.com/nicollasm/golang-postgre-importer-csv/pkg"
)

const (
	DB_CONFIG_FILE = "./dbconfig.json"
)

func main() {
	start := time.Now()

	db, err := pkg.InitDB(DB_CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pkg.ReadAndWriteToDB(db, "./example.csv", "myTable")

	elapsed := time.Since(start)
	fmt.Printf("\nTempo total de execução: %s\n", elapsed)
}
