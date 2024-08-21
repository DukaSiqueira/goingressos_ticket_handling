package services

import (
	"goingressos_ticket_handling/handlers"
	"goingressos_ticket_handling/utils"
)

var logger = utils.NewLogger()

func RunStatusAndIntegrityCheck() error {
	for {
		buy, err := handlers.FetchNextBuyByStatus()
		if err != nil {
			logger.Fatal("Erro ao buscar próxima compra: ", err)
			return err
		}
		if buy == nil {
			logger.Info("Nenhuma compra restante para processar.")
			break
		}

		logger.Info("Processando compra ID:", buy.ID, "Status:", buy.Status, "EventID:", buy.EventID)

		err = handlers.UpdateRecordsAndTicketsForBuy(*buy)
		if err != nil {
			logger.Fatal("Erro ao atualizar registros para Buy ID:", buy.ID, err)
			return err
		}

		logger.Info("Compra ID:", buy.ID, "processada com sucesso.")
	}

	logger.Info("Verificação de status e integridade realizada com sucesso.")
	return nil
}
