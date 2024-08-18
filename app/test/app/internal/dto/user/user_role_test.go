package user

import (
	"golang-template-api-service/app/internal/dto/user"
	"testing"
)

func TestUserRole(t *testing.T) {
	// Create a sample UserRole instance
	role := user.NewUserRole("1", "Operator")

	// Test the GetID method
	expectedID := "1"
	actualID := role.GetID()
	if actualID != expectedID {
		t.Errorf("GetID() returned %s, expected %s", actualID, expectedID)
	}

	// Test the GetName method
	expectedName := "Operator"
	actualName := role.GetName()
	if actualName != expectedName {
		t.Errorf("GetName() returned %s, expected %s", actualName, expectedName)
	}
}
