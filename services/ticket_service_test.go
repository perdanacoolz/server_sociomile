package services_test

import (
	"testing"

	"github.com/perdana/sociomile/models"
	"github.com/perdana/sociomile/repositories"
	"github.com/perdana/sociomile/services"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestTicketService_UpdateStatus(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.Ticket{}, &models.TicketEvent{})

	repo := repositories.NewTicketRepo(db, nil) // no redis for test
	service := services.NewTicketService(repo)

	ticket := &models.Ticket{ID: 1, TenantID: 1, Status: "open"}
	db.Create(ticket)

	err := service.UpdateStatus(1, "resolved", 1)
	assert.NoError(t, err)

	var event models.TicketEvent
	db.First(&event)
	assert.Equal(t, "status_changed_to_resolved", event.Event)
}
