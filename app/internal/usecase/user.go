package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-template-api-service/app/config"
	"golang-template-api-service/app/internal/dto/common"
	"golang-template-api-service/app/internal/dto/user"
	"golang-template-api-service/app/internal/entity"
	"golang-template-api-service/app/internal/repository"
	"golang-template-api-service/app/utils"
	"strings"

	"github.com/google/uuid"
)

type UserUsecase struct {
	UserRepository *repository.UserRepository
	httpClient     *utils.HTTPClient
}

func NewUserUsecase(UserRepository *repository.UserRepository, httpClient *utils.HTTPClient) *UserUsecase {
	return &UserUsecase{UserRepository: UserRepository,
		httpClient: httpClient,
	}
}

func (u *UserUsecase) CreateUser(user user.UserRequest) (*entity.User, error) {

	emailLower := strings.ToLower(user.Email)

	newUser := &entity.User{
		Role:     &user.UserRole,
		UserType: &user.UserType,
		Email:    &emailLower,
	}
	return u.UserRepository.CreateUser(newUser)
}

func (u *UserUsecase) GetUserDetails(userId uuid.UUID, token string) (*user.UserDetailResponse, error) {

	services := config.LoadViperConfig().Services

	// Set the base URL for this particular use case
	u.httpClient.SetBaseURL(services.TestServiceUrl)

	// Set the access token in the httpClient's request headers
	u.httpClient.SetHeader("Authorization", token)
	url := fmt.Sprintf("/users/%s/detail", userId.String())

	// Perform the HTTP GET request using the httpClient
	response, err := u.httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	var userDetails user.UserDetailResponse
	var responseMap common.Response
	err = json.Unmarshal([]byte(response), &responseMap)
	if err != nil {
		return nil, err
	}
	userDetailJSON, err := json.Marshal(responseMap.Data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	err = json.Unmarshal(userDetailJSON, &userDetails)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err

	}

	user, err := u.UserRepository.GetUserByEmail(userDetails.Email)
	fmt.Printf("GetUserByEmail  %s\n", user)

	if user == nil {
		return nil, err
	}
	if user.UserRegisterId != userDetails.UserId {
		user.UserRegisterId = userDetails.UserId
		u.UserRepository.UpdateUser(user)
	}
	profilePictureResponse, err := utils.GeneratePresignedURL(*user.ProfilePicture)

	if err != nil {
		userDetails.ProfilePicture = ""
	} else {
		userDetails.ProfilePicture = profilePictureResponse.ObjectURL
	}
	return &userDetails, nil
}

func (u *UserUsecase) DeleteUser(userId uuid.UUID, token string) (bool, error) {

	user, err := u.UserRepository.GetUserById(userId)
	fmt.Printf("GetUserByEmail  %s\n", user)
	if err != nil {
		return false, err
	}

	*user.IsDeleted = true
	u.UserRepository.UpdateUser(user)

	return true, nil
}
func (u *UserUsecase) UpdateUser(userId uuid.UUID, userStatus user.UserStatusUpdate, token string) (bool, error) {

	services := config.LoadViperConfig().Services

	// Set the base URL for this particular use case
	u.httpClient.SetBaseURL(services.TestServiceUrl)

	// Set the access token in the httpClient's request headers
	u.httpClient.SetHeader("Authorization", token)
	url := fmt.Sprintf("/users/%s/detail", userId.String())

	// Perform the HTTP GET request using the httpClient
	response, err := u.httpClient.Get(url)

	if err != nil {
		return false, err
	}

	var userDetails user.UserDetailResponse
	var responseMap common.Response
	err = json.Unmarshal([]byte(response), &responseMap)
	if err != nil {
		return false, err
	}
	userDetailJSON, err := json.Marshal(responseMap.Data)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}

	err = json.Unmarshal(userDetailJSON, &userDetails)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err

	}

	soraUser, err := u.UserRepository.GetUserByEmail(userDetails.Email)
	fmt.Printf("GetUserByEmail  %s\n", soraUser)

	if soraUser == nil {
		return false, err
	}

	userEnable := &user.UserEnableRequest{
		FirstName: userStatus.FirstName,
		LastName:  userStatus.LastName,
		Email:     userDetails.Email,
		Phone:     userStatus.Phone,
		UserID:    userDetails.UserId,
		Roles:     transformToRoleNames(userDetails),
		Country:   userStatus.Country,
		IsActive:  userStatus.Active,
	}

	// Set the base URL for this particular use case
	u.httpClient.SetBaseURL(services.TestServiceUrl)
	u.httpClient.SetHeader("Authorization", token)
	u.httpClient.SetHeader("Content-Type", "application/json")

	updateUrl := fmt.Sprintf("/users/%s", userId.String())
	// Convert user struct to JSON
	userJSON, err := json.Marshal(userEnable)
	if err != nil {
		return false, err
	}

	// Create a bytes.Buffer from the JSON data
	userBuffer := bytes.NewBuffer(userJSON)
	// Make a POST request
	updateResponse, err := u.httpClient.Put(updateUrl, userBuffer)
	if err != nil {
		return false, err
	}
	fmt.Printf("Update Response  %s\n", updateResponse)

	*soraUser.IsActive = userStatus.Active
	u.UserRepository.UpdateUser(soraUser)
	if err != nil {
		return false, err
	}
	return true, nil
}

func transformToRoleNames(user user.UserDetailResponse) []string {
	roleNames := make([]string, len(user.UserRole))
	for i, role := range user.UserRole {
		roleNames[i] = role.Name
	}
	return roleNames
}
