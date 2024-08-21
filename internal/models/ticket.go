package models

import "time"

type Ticket struct {
	ID        uint   `gorm:"primaryKey"`
	EventID   uint   `gorm:"column:eventId;not null"`
	UserID    uint   `gorm:"not null"`
	Status    string `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
