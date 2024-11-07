package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vicsimopher/launch-scheduler/internal/handlers"
	"github.com/vicsimopher/launch-scheduler/internal/repository"
	"github.com/vicsimopher/launch-scheduler/internal/service"
)

func main() {
	// Initialize DB
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize repositories, services, and handlers
	bookingRepo := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepo)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	// Setup Gin router
	r := gin.Default()

	// Routes
	r.POST("/bookings", bookingHandler.CreateBooking)
	r.GET("/bookings", bookingHandler.GetBookings)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
