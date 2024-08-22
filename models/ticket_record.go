package models

import "gorm.io/gorm"

type TicketRecord struct {
	gorm.Model
	CompID      uint64 `gorm:"column:compId"`
	TicketID    uint64 `gorm:"column:ticketId"`
	EventID     uint64 `gorm:"column:eventId"`
	TicketBuyID uint64 `gorm:"column:ticketBuyId"`
	BatchID     uint64 `gorm:"column:batchId"`
	Status      int
	Price       float64
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (TicketRecord) TableName() string {
	return "tickets_records"
}
