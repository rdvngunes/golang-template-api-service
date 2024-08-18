package enums

// Define a custom type for the enum
type UserRole int

// Define enum values using constants and iota
const (
	Admin UserRole = iota
	Regular
)

func (s UserRole) String() string {
	switch s {
	case Admin:
		return "Admin"
	case Regular:
		return "Regular"
	default:
		return "Unknown"
	}
}
