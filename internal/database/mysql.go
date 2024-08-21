package database

import (
	"database/sql"
	"fmt"
	"log"

	"goingressos_ticket_handling/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
