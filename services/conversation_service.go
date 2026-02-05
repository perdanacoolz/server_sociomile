package services

import (
	"github.com/perdana/sociomile/models"
	"github.com/perdana/sociomile/repositories"
)

type ConversationService struct {
	repo *repositories.MessageRepo
}

func NewConversationService(repo *repositories.MessageRepo) *ConversationService {
	return &ConversationService{repo}
}

func (s *ConversationService) Send(message *models.Message) error {
	return s.repo.Create(message)
}

func (s *ConversationService) Get(ticketID uint) ([]models.Message, error) {
	return s.repo.GetByTicketID(ticketID)
}
