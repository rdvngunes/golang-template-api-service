package user

type UserRequest struct {
	UserRole string `json:"user_role" binding:"required"`
	UserType string `json:"user_type" binding:"required"`
	Email    string `json:"email"  binding:"required"`
	ReOC     string `json:"reoc"`
}
