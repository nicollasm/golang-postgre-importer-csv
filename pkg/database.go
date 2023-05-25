package pkg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// Inicialização do banco de dados, retorna erro se falhar
func InitDB(cfg DBConfig) (*sql.DB, error) {
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
