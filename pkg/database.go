package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func InitDB(configFile string) (*sql.DB, error) {
	var cfg DBConfig

	file, err := os.Open(configFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println("Conectado ao banco de dados com sucesso.")
	return db, nil
}
