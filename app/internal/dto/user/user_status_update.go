package user

type UserStatusUpdate struct {
	FirstName string `json:"first_name"`
	Country   string `json:"country"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Active    bool   `json:"active"` // true for activation, false for deactivation
}
