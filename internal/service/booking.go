package service

import (
	"github.com/vicsimopher/launch-scheduler/internal/models"
	"github.com/vicsimopher/launch-scheduler/internal/repository"
)

type BookingService interface {
	CreateBooking(request *models.BookingRequest) error
	GetBookings() ([]models.Booking, error)
}

type bookingService struct {
	repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{repo: repo}
}

func (s *bookingService) CreateBooking(request *models.BookingRequest) error {
	booking := &models.Booking{
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Gender:        request.Gender,
		Birthday:      request.Birthday,
		LaunchpadID:   request.LaunchpadID,
		DestinationID: request.DestinationID,
		LaunchDate:    request.LaunchDate,
	}

	// TODO: Add SpaceX API integration to check launch pad availability

	return s.repo.CreateBooking(booking)
}

func (s *bookingService) GetBookings() ([]models.Booking, error) {
	return s.repo.GetBookings()
}
