package user

type UserRole struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewUserRole(id, name string) *UserRole {
	return &UserRole{id, name}
}

// GetID returns the ID field value of UserRole.
func (ur *UserRole) GetID() string {
	return ur.Id
}

// GetName returns the Name field value of UserRole.
func (ur *UserRole) GetName() string {
	return ur.Name
}
