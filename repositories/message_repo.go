package repositories

import (
	"github.com/perdana/sociomile/models"

	"gorm.io/gorm"
)

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) *MessageRepo {
	return &MessageRepo{db}
}

func (r *MessageRepo) Create(message *models.Message) error {
	return r.db.Create(message).Error
}

func (r *MessageRepo) GetByTicketID(ticketID uint) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Where("ticket_id = ?", ticketID).Order("created_at ASC").Find(&messages).Error
	return messages, err
}
