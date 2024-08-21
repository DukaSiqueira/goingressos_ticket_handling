package services

import (
	"goingressos_ticket_handling/config"
	"goingressos_ticket_handling/models"
)

// var logger = utils.NewLogger()

// Verifica e corrige inconsistências entre as tabelas buys, ticket_records e tickets
func RunIntegrityCheck() error {
	// Corrigir o campo 'sold' da tabela tickets baseado nas vendas concluídas
	err := fixSoldCount()
	if err != nil {
		return err
	}

	// Corrigir status inconsistentes entre buys e ticket_records
	err = fixStatusInconsistencies()
	if err != nil {
		return err
	}

	return nil
}

// Corrige o valor de 'sold' na tabela tickets baseado no número de vendas aprovadas
func fixSoldCount() error {
	var tickets []models.Ticket
	err := config.DB.Find(&tickets).Error
	if err != nil {
		logger.Fatal("Erro ao buscar tickets:", err)
		return err
	}

	for _, ticket := range tickets {
		var totalSold int64
		err := config.DB.Model(&models.TicketRecord{}).Where("ticket_id = ? AND status = 1", ticket.ID).Count(&totalSold).Error
		if err != nil {
			logger.Fatal("Erro ao contar tickets vendidos:", err)
			return err
		}

		if int(totalSold) != ticket.Sold {
			ticket.Sold = int(totalSold)
			err := config.DB.Save(&ticket).Error
			if err != nil {
				logger.Fatal("Erro ao atualizar sold no ticket:", err)
				return err
			}
		}
	}

	return nil
}

// Corrige status inconsistentes entre buys e ticket_records
func fixStatusInconsistencies() error {
	var records []models.TicketRecord
	result := config.DB.Preload("Buy").Find(&records)
	if result.Error != nil {
		logger.Fatal("Erro ao buscar registros de tickets:", result.Error)
		return result.Error
	}

	for _, record := range records {
		if record.Status != record.Buy.Status {
			record.Status = record.Buy.Status
			err := config.DB.Save(&record).Error
			if err != nil {
				logger.Fatal("Erro ao corrigir status no ticket record:", err)
				return err
			}
		}
	}

	return nil
}
