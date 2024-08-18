package enums

import (
	"golang-template-api-service/app/internal/dto/enums"
	"testing"
)

func TestUserRoleString(t *testing.T) {
	tests := []struct {
		role     enums.UserRole
		expected string
	}{
		{enums.Admin, "Admin"},
		{enums.Regular, "Regular"},
		{enums.UserRole(-1), "Unknown"},
	}

	for _, test := range tests {
		result := test.role.String()
		if result != test.expected {
			t.Errorf("UserRole(%d).String() returned %s; expected %s", test.role, result, test.expected)
		}
	}
}
