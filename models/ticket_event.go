package models

import "time"

type TicketEvent struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	TicketID  uint      `json:"ticket_id"`
	Event     string    `json:"event"` // e.g., "status_changed_to_resolved"
	CreatedAt time.Time `json:"created_at"`
}
