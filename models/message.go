package models

import "time"

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	TicketID  uint      `json:"ticket_id" gorm:"index"`
	SenderID  uint      `json:"sender_id"` // user ID (agent/customer)
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at" gorm:"index"`
}
