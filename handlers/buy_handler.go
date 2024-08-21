package handlers

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/models"
)

func GetBuysByStatus(status int) ([]models.Buy, error) {
	var buys []models.Buy
	result := config.DB.Where("status = ?", status).Find(&buys)
	return buys, result.Error
}
