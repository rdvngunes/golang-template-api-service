package user

import "github.com/google/uuid"

type UserEnableRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	AnspID    string    `json:"ansp_id"`
	UserID    uuid.UUID `json:"user_id"`
	Roles     []string  `json:"roles"`
	Country   string    `json:"country"`
	IsActive  bool      `json:"is_active"`
}
