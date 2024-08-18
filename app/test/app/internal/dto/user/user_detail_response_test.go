package user

import (
	"golang-template-api-service/app/internal/dto/user"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserDetailResponse(t *testing.T) {
	// Setup test data
	userID := uuid.New()
	userRoles := []user.UserRole{
		{
			Id:   uuid.New().String(),
			Name: "Admin",
		},
		{
			Id:   uuid.New().String(),
			Name: "User",
		},
	}
	createdDate := time.Now()

	// Create an instance of UserDetailResponse
	userDetail := user.UserDetailResponse{
		UserId:         userID,
		UserName:       "johndoe",
		IsActive:       true,
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john.doe@example.com",
		Country:        "USA",
		Phone:          "+123456789",
		UserRole:       userRoles,
		ProfilePicture: "http://example.com/profile.jpg",
		UserType:       "TypeA",
		CreatedDate:    createdDate,
	}

	// Assertions
	assert.Equal(t, userID, userDetail.UserId, "UserId should match")
	assert.Equal(t, "johndoe", userDetail.UserName, "UserName should match")
	assert.Equal(t, true, userDetail.IsActive, "IsActive should be true")
	assert.Equal(t, "John", userDetail.FirstName, "FirstName should match")
	assert.Equal(t, "Doe", userDetail.LastName, "LastName should match")
	assert.Equal(t, "john.doe@example.com", userDetail.Email, "Email should match")
	assert.Equal(t, "USA", userDetail.Country, "Country should match")
	assert.Equal(t, "+123456789", userDetail.Phone, "Phone should match")
	assert.Equal(t, userRoles, userDetail.UserRole, "UserRole should match")
	assert.Equal(t, "http://example.com/profile.jpg", userDetail.ProfilePicture, "ProfilePicture should match")
	assert.Equal(t, "TypeA", userDetail.UserType, "UserType should match")
	assert.WithinDuration(t, createdDate, userDetail.CreatedDate, time.Second, "CreatedDate should match within a second")
}
