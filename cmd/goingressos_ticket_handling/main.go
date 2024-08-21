package main

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/internal/database"
	"goingressos_ticket_handling/internal/tickets"
	"goingressos_ticket_handling/pkg/logger"
)

func main() {
	logger := logger.NewLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Erro ao carregar as configurações:", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		logger.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	tickets.ProcessTicketRefunds(db)

	// Exemplo de loop para execução contínua
	// for {
	//     tickets.ProcessTicketRefunds(db)
	//     time.Sleep(5 * time.Minute)
	// }
}
