package auth

import (
	"golang-template-api-service/app/internal/dto/auth"
	"testing"
)

func TestRoles(t *testing.T) {
	// Create an instance of Roles
	roles := auth.Roles{
		Roles: []string{"role1", "role2"},
	}

	// Check the fields of the roles
	if len(roles.Roles) != 2 {
		t.Error("Roles should contain 2 roles")
	}

	if roles.Roles[0] != "role1" || roles.Roles[1] != "role2" {
		t.Error("Roles should contain 'role1' and 'role2'")
	}
}
