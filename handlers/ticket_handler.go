package handlers

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/models"
)

func UpdateTicketSoldCount(eventID, ticketID uint64) error {
	var ticket models.Ticket

	result := config.DB.
		Where("eventId = ? AND id = ?", eventID, ticketID).
		First(&ticket)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			logger.Info("Nenhum ticket encontrado para o evento", eventID, "e ticket", ticketID)
			return nil
		}
		logger.Fatal("Erro ao buscar ticket para o evento:", eventID, "e ticket:", ticketID, result.Error)
		return result.Error
	}

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

	return nil
}
