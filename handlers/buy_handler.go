package handlers

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/models"
)

func FetchNextBuyByStatus() (*models.Buy, error) {
	var buy models.Buy
	statusList := []int{99, 98, 14, 9, 4, 3, 2, 0}

	result := config.DB.
		Where("status IN ?", statusList).
		Order("created_at ASC").
		First(&buy)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, nil // Não há mais compras para processar
		}
		return nil, result.Error
	}

	return &buy, nil
}
