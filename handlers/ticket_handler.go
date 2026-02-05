package handlers

import (
	"net/http"

	"github.com/perdana/sociomile/config"
	"github.com/perdana/sociomile/models"
	"github.com/perdana/sociomile/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TicketHandler struct {
	service *services.TicketService
}

func NewTicketHandler(service *services.TicketService) *TicketHandler {
	return &TicketHandler{service}
}

func (h *TicketHandler) Create(c *gin.Context) {
	var ticket models.Ticket
	if err := c.BindJSON(&ticket); err != nil {
		config.Logger.Error("bind error", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket.TenantID = uint(c.GetFloat64("tenant_id"))
	if err := h.service.Create(&ticket); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ticket)
}
