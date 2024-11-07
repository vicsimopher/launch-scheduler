package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicsimopher/launch-scheduler/internal/models"
	"github.com/vicsimopher/launch-scheduler/internal/service"
)

type BookingHandler struct {
	service service.BookingService
}

func NewBookingHandler(service service.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var request models.BookingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateBooking(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking created successfully"})
}

func (h *BookingHandler) GetBookings(c *gin.Context) {
	bookings, err := h.service.GetBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}
