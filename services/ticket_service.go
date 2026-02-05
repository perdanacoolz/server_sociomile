package services

import (
	"github.com/perdana/sociomile/models"
	"github.com/perdana/sociomile/repositories"
)

type TicketService struct {
	repo *repositories.TicketRepo
}

func NewTicketService(repo *repositories.TicketRepo) *TicketService {
	return &TicketService{repo}
}

func (s *TicketService) Create(ticket *models.Ticket) error {
	return s.repo.Create(ticket)
}

func (s *TicketService) Assign(ticketID uint, agentID uint, tenantID uint) error {
	return s.repo.Assign(ticketID, agentID, tenantID)
}

func (s *TicketService) UpdateStatus(ticketID uint, status string, tenantID uint) error {
	err := s.repo.UpdateStatus(ticketID, status, tenantID)
	if err == nil && status == "resolved" {
		event := &models.TicketEvent{TicketID: ticketID, Event: "status_changed_to_" + status}
		s.repo.LogEvent(event) // Sync for simplicity; could be async with goroutine
	}
	return err
}

func (s *TicketService) List(tenantID uint, status, agentID *string) ([]models.Ticket, error) {
	return s.repo.List(tenantID, status, agentID)
}
