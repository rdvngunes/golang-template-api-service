package user

import (
	"time"

	"github.com/google/uuid"
)

type UserDetailResponse struct {
	UserId           uuid.UUID  `json:"user_id"`
	UserName         string     `json:"user_name"`
	IsActive         bool       `json:"is_active"`
	FirstName        string     `json:"first_name"`
	LastName         string     `json:"last_name"`
	Email            string     `json:"email"`
	Country          string     `json:"country"`
	Phone            string     `json:"phone"`
	UserRole         []UserRole `json:"user_roles"`
	AnspID           string     `json:"ansp_id"`
	OrganizationID   string     `json:"organization_id"`
	OrganizationName string     `json:"organization_name"`
	ProfilePicture   string     `json:"profile_picture"`
	UserSoraRole     string     `json:"user_sora_role"`
	ReOC             string     `json:"reoc"`
	UserType         string     `json:"user_type"`
	CreatedDate      time.Time  `json:"created_date"`
}
