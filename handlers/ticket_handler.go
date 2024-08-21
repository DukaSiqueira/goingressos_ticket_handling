package handlers

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/models"
)

func UpdateTicketSoldCount(eventID, batchID uint64) error {
	var tickets []models.Ticket

	result := config.DB.
		Where("eventId = ? AND batchId = ?", eventID, batchID).
		Find(&tickets)
	if result.Error != nil {
		logger.Fatal("Erro ao buscar tickets para o evento:", eventID, "e batch:", batchID, result.Error)
		return result.Error
	}

	for _, ticket := range tickets {
		var totalSold int64
		config.DB.Model(&models.TicketRecord{}).
			Where("ticketId = ? AND status = 1", ticket.ID).
			Count(&totalSold)

		ticket.Sold = int(totalSold)
		err := config.DB.Save(&ticket).Error
		if err != nil {
			logger.Fatal("Erro ao atualizar o campo 'sold' para o ticket:", ticket.ID, err)
			return err
		}

		logger.Info("Campo 'sold' atualizado para Ticket ID:", ticket.ID, "com sucesso. Total vendido:", totalSold)
	}

	return nil
}
