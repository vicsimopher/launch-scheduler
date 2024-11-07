package models

import "time"

type Booking struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	FirstName     string    `json:"first_name" binding:"required"`
	LastName      string    `json:"last_name" binding:"required"`
	Gender        string    `json:"gender" binding:"required"`
	Birthday      time.Time `json:"birthday" binding:"required"`
	LaunchpadID   string    `json:"launchpad_id" binding:"required"`
	DestinationID string    `json:"destination_id" binding:"required"`
	LaunchDate    time.Time `json:"launch_date" binding:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type BookingRequest struct {
	FirstName     string    `json:"first_name" binding:"required"`
	LastName      string    `json:"last_name" binding:"required"`
	Gender        string    `json:"gender" binding:"required"`
	Birthday      time.Time `json:"birthday" binding:"required"`
	LaunchpadID   string    `json:"launchpad_id" binding:"required"`
	DestinationID string    `json:"destination_id" binding:"required"`
	LaunchDate    time.Time `json:"launch_date" binding:"required"`
}
