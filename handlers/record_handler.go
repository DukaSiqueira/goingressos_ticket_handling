package handlers

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/models"
	"goingressos_ticket_handling/utils"
)

var logger = utils.NewLogger()

func UpdateRecordsAndTicketsForBuy(buy models.Buy) error {
	var ticketRecords []models.TicketRecord
	result := config.DB.
		Where("ticketBuyId = ? AND eventId = ?", buy.ID, buy.EventID).
		Find(&ticketRecords)

	if result.Error != nil {
		logger.Fatal("Erro ao buscar registros de tickets para Buy ID:", buy.ID, result.Error)
		return result.Error
	}

	logger.Info("Encontrados", len(ticketRecords), "registros de tickets para Buy ID:", buy.ID)

	for _, record := range ticketRecords {
		logger.Info("Atualizando TicketRecord ID:", record.ID, "com status:", buy.Status)

		record.Status = buy.Status
		err := config.DB.Save(&record).Error
		if err != nil {
			logger.Fatal("Erro ao atualizar TicketRecord ID:", record.ID, err)
			return err
		}

		logger.Info("TicketRecord ID:", record.ID, "atualizado com sucesso.")
	}

	err := UpdateTicketSoldCount(buy.EventID, ticketRecords[0].BatchID)
	if err != nil {
		logger.Fatal("Erro ao atualizar o campo 'sold' para o evento:", buy.EventID, err)
		return err
	}

	return nil
}
