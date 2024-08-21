package handlers

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/models"
)

func UpdateTicketSold(ticketID uint64, sold int) error {
	result := config.DB.Model(&models.Ticket{}).Where("id = ?", ticketID).Update("sold", sold)
	return result.Error
}
