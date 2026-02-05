package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/perdana/sociomile/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type TicketRepo struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewTicketRepo(db *gorm.DB, redis *redis.Client) *TicketRepo {
	return &TicketRepo{db, redis}
}

func (r *TicketRepo) Create(ticket *models.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *TicketRepo) Assign(ticketID uint, agentID uint, tenantID uint) error {
	return r.db.Model(&models.Ticket{}).Where("id = ? AND tenant_id = ?", ticketID, tenantID).Update("assigned_agent_id", agentID).Error
}

func (r *TicketRepo) UpdateStatus(ticketID uint, status string, tenantID uint) error {
	return r.db.Model(&models.Ticket{}).Where("id = ? AND tenant_id = ?", ticketID, tenantID).Update("status", status).Error
}

func (r *TicketRepo) List(tenantID uint, status, agentID *string) ([]models.Ticket, error) {
	cacheKey := fmt.Sprintf("tickets:%d:%s:%s", tenantID, *status, *agentID)
	cached, err := r.redis.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var tickets []models.Ticket
		json.Unmarshal([]byte(cached), &tickets)
		return tickets, nil
	}

	query := r.db.Where("tenant_id = ?", tenantID)
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if agentID != nil {
		query = query.Where("assigned_agent_id = ?", *agentID)
	}
	var tickets []models.Ticket
	err = query.Find(&tickets).Error
	if err == nil {
		data, _ := json.Marshal(tickets)
		r.redis.Set(context.Background(), cacheKey, data, 5*time.Minute)
	}
	return tickets, err
}

func (r *TicketRepo) LogEvent(event *models.TicketEvent) error {
	return r.db.Create(event).Error
}
