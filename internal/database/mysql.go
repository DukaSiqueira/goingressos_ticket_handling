package database

import (
	"goingressos_ticket_handling/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}

	return db, nil
}
