package models

import "time"

type Ticket struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`   // open, in_progress, resolved, closed
	Priority        string    `json:"priority"` // low, medium, high
	AssignedAgentID uint      `json:"assigned_agent_id"`
	CustomerID      uint      `json:"customer_id"`
	TenantID        uint      `json:"tenant_id" gorm:"index"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
