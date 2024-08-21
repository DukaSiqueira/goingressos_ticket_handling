package models

import (
	"gorm.io/gorm"
)

type TicketRecord struct {
	gorm.Model
	TicketID uint64 `gorm:"column:ticketId"`
	EventID  uint64 `gorm:"column:eventId"`
	Status   int
	Price    float64
	BuyID    uint64 `gorm:"column:buyId"`
	Buy      Buy    `gorm:"foreignKey:BuyID"`
}
