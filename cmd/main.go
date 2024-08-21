package main

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/services"
	"goingressos_ticket_handling/utils"
	"time"
)

func main() {
	logger := utils.NewLogger()

	config.LoadConfig()

	ticker := time.NewTicker(3 * time.Hour)
	defer ticker.Stop()

	// for {
	// 	select {
	// 	case <-ticker.C:
	logger.Info("Iniciando verificação de status e integridade.")
	err := services.RunStatusCheck()
	if err != nil {
		logger.Fatal("Erro ao verificar status e integridade:", err)
	} else {
		logger.Info("Verificação de status e integridade concluída.")
	}
	// 	}
	// }
}
