package handlers

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/models"
	"time"
)

func FetchNextBuyByStatus() (*models.Buy, error) {
	var buy models.Buy
	statusList := []int{99, 98, 14, 0}

	result := config.DB.
		Where("status IN ?", statusList).
		Where("created_at < ?", time.Now().Add(-3*time.Hour)).
		Order("created_at ASC").
		First(&buy)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, nil
		}
		return nil, result.Error
	}

	return &buy, nil
}
