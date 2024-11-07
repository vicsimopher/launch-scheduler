package repository

import (
	"fmt"
	"os"

	"github.com/vicsimopher/launch-scheduler/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *models.Booking) error
	GetBookings() ([]models.Booking, error)
}

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&models.Booking{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) CreateBooking(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

func (r *postgresRepository) GetBookings() ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Find(&bookings).Error
	return bookings, err
}
