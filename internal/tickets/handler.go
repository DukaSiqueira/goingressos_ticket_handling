package tickets

import (
	"goingressos_ticket_handling/internal/models"
	"time"

	"gorm.io/gorm"
)

func ProcessTicketRefunds(db *gorm.DB) {
	var tickets []models.Ticket

	db.Where("status = ?", "pending_refund").Find(&tickets)

	for _, ticket := range tickets {
		ticket.Status = "refunded"
		ticket.UpdatedAt = time.Now()

		db.Save(&ticket)
	}
}
