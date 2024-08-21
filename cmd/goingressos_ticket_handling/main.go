package main

import (
	"os"
	"os/signal"
	"syscall"

	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/internal/tickets"
	"goingressos_ticket_handling/pkg/logger"
)

func main() {
	logger := logger.NewLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Erro ao carregar configurações:", err)
	}
	logger.Info("Iniciando serviço de devolução de ingressos...")
	tickets.StartTicketHandler(cfg, logger)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	logger.Info("Encerrando serviço...")
}
