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

	for {
		logger.Info("Iniciando processo de verificação de status e integridade.")
		err := services.RunStatusAndIntegrityCheck()
		if err != nil {
			logger.Fatal("Erro ao executar a verificação: ", err)
		} else {
			logger.Info("Verificação de status e integridade concluída.")
		}

		time.Sleep(5 * time.Second)
	}
}
