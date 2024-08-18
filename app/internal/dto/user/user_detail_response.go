package user

import (
	"time"

	"github.com/google/uuid"
)

type UserDetailResponse struct {
	UserId         uuid.UUID  `json:"user_id"`
	UserName       string     `json:"user_name"`
	IsActive       bool       `json:"is_active"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Email          string     `json:"email"`
	Country        string     `json:"country"`
	Phone          string     `json:"phone"`
	UserRole       []UserRole `json:"user_roles"`
	ProfilePicture string     `json:"profile_picture"`
	UserType       string     `json:"user_type"`
	CreatedDate    time.Time  `json:"created_date"`
}
