package config

import (
	"goingressos_ticket_handling/utils"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

var DB *gorm.DB

func LoadConfig() {
	logger := utils.NewLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Erro ao carregar o arquivo .env: ", err)
	}

	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Erro ao conectar no banco de dados: ", err)
	}

	DB = db
	logger.Info("Banco de dados conectado com sucesso!!")
}
