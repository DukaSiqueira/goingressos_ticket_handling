package tickets

import (
	"database/sql"
	"time"

	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/internal/database"
	"goingressos_ticket_handling/pkg/logger"
)

func StartTicketHandler(cfg *config.Config, logger *logger.Logger) {
	db, err := database.Connect(cfg)
	if err != nil {
		logger.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	go func() {
		for {
			processTickets(db, logger)

			time.Sleep(5 * time.Minute)
		}
	}()
}

func processTickets(_ *sql.DB, logger *logger.Logger) {
	logger.Info("Processando devoluções de ingressos...")
	logger.Info("Encerrando processamento de devoluções de ingressos")
}
