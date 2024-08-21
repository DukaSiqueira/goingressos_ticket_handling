package handlers

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/models"
)

func GetRecordsByStatus(status int) ([]models.TicketRecord, error) {
	var records []models.TicketRecord
	result := config.DB.Where("status = ?", status).Find(&records)
	return records, result.Error
}
