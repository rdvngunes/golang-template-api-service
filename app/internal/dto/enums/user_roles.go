package enums

// Define a custom type for the enum
type UserRole int

// Define enum values using constants and iota
const (
	Operator UserRole = iota
	Reviever
)

func (s UserRole) String() string {
	switch s {
	case Operator:
		return "Operator"
	case Reviever:
		return "Reviever"
	default:
		return "Unknown"
	}
}
