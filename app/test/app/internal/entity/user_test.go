package entity

import (
	"golang-template-api-service/app/internal/entity"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserStructInitialization(t *testing.T) {
	// Setup test data
	userID := uuid.New()
	userRegisterID := uuid.New()
	email := "john.doe@example.com"
	role := "Admin"
	userType := "TypeA"
	profilePicture := "http://example.com/profile.jpg"
	isActive := true
	isDeleted := false
	createdOn := time.Now()
	modifiedOn := time.Now()

	// Create an instance of User
	userData := entity.User{
		UserId:         userID,
		UserRegisterId: userRegisterID,
		Email:          &email,
		Role:           &role,
		UserType:       &userType,
		ProfilePicture: &profilePicture,
		IsActive:       &isActive,
		IsDeleted:      &isDeleted,
		CreatedOn:      createdOn,
		ModifiedOn:     modifiedOn,
	}

	// Assertions
	assert.Equal(t, userID, userData.UserId, "UserId should match")
	assert.Equal(t, userRegisterID, userData.UserRegisterId, "UserRegisterId should match")
	assert.Equal(t, &email, userData.Email, "Email should match")
	assert.Equal(t, &role, userData.Role, "Role should match")
	assert.Equal(t, &userType, userData.UserType, "UserType should match")
	assert.Equal(t, &profilePicture, userData.ProfilePicture, "ProfilePicture should match")
	assert.Equal(t, &isActive, userData.IsActive, "IsActive should be true")
	assert.Equal(t, &isDeleted, userData.IsDeleted, "IsDeleted should be false")
	assert.WithinDuration(t, createdOn, userData.CreatedOn, time.Second, "CreatedOn should match within a second")
	assert.WithinDuration(t, modifiedOn, userData.ModifiedOn, time.Second, "ModifiedOn should match within a second")
}

func TestUserDefaultValues(t *testing.T) {
	// Create an instance of User with default values
	userData := entity.User{}

	// Assertions
	assert.NotEqual(t, uuid.Nil, userData.UserId, "UserId should be generated and not be nil")
	assert.NotEqual(t, uuid.Nil, userData.UserRegisterId, "UserRegisterId should be generated and not be nil")
	assert.Nil(t, userData.Email, "Email should be nil by default")
	assert.Nil(t, userData.Role, "Role should be nil by default")
	assert.Nil(t, userData.UserType, "UserType should be nil by default")
	assert.Nil(t, userData.ProfilePicture, "ProfilePicture should be nil by default")
	assert.Nil(t, userData.IsActive, "IsActive should be nil by default")
	assert.Nil(t, userData.IsDeleted, "IsDeleted should be nil by default")
	assert.WithinDuration(t, time.Now(), userData.CreatedOn, time.Second, "CreatedOn should be set to current time by default")
	assert.WithinDuration(t, time.Now(), userData.ModifiedOn, time.Second, "ModifiedOn should be set to current time by default")
}
