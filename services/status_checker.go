package services

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/models"
	"goingressos_ticket_handling/utils"
)

var logger = utils.NewLogger()

func RunStatusCheck() error {
	err := fixExpiredBuysAndRecords()
	if err != nil {
		return err
	}

	return nil
}

func fixExpiredBuysAndRecords() error {
	statusList := []int{98, 99, 0, 3, 4, 5, 8, 9, 14, 171}
	var buys []models.Buy
	result := config.DB.Where("status IN ?", statusList).Find(&buys)

	if result.Error != nil {
		logger.Fatal("Erro ao buscar buys expirados:", result.Error)
		return result.Error
	}

	for _, buy := range buys {
		buy.Status = 4
		err := config.DB.Save(&buy).Error
		if err != nil {
			logger.Fatal("Erro ao atualizar buy:", err)
			return err
		}
	}

	return nil
}
