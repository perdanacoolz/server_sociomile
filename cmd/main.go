package main

import (
	"github.com/perdana/sociomile/config"
	"github.com/perdana/sociomile/handlers"
	"github.com/perdana/sociomile/middleware"
	"github.com/perdana/sociomile/models"
	"github.com/perdana/sociomile/repositories"
	"github.com/perdana/sociomile/services"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config.Init()
	cfg := config.Load()

	db, err := gorm.Open(mysql.Open(cfg.DBDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Tenant{}, &models.User{}, &models.Ticket{}, &models.Message{}, &models.TicketEvent{})

	redisClient := redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})

	userRepo := repositories.NewUserRepo(db)
	ticketRepo := repositories.NewTicketRepo(db, redisClient)
	messageRepo := repositories.NewMessageRepo(db)

	authService := services.NewAuthService(userRepo)
	ticketService := services.NewTicketService(ticketRepo)
	conversationService := services.NewConversationService(messageRepo)

	authHandler := handlers.NewAuthHandler(authService)
	ticketHandler := handlers.NewTicketHandler(ticketService)
	conversationHandler := handlers.NewConversationHandler(conversationService) // Implement similar

	r := gin.Default()

	r.POST("/login", authHandler.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware("admin", "agent", "customer"))

	adminAgent := r.Group("/")
	adminAgent.Use(middleware.AuthMiddleware("admin", "agent"))

	// Tickets
	adminAgent.POST("/tickets", ticketHandler.Create)
	adminAgent.PUT("/tickets/:id/assign", ticketHandler.Assign)
	adminAgent.PUT("/tickets/:id/status", ticketHandler.UpdateStatus)
	protected.GET("/tickets", ticketHandler.List)

	// Conversations
	protected.POST("/conversations/:ticket_id", conversationHandler.Send)
	protected.GET("/conversations/:ticket_id", conversationHandler.Get)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
