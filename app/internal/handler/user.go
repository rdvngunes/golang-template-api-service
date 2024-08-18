package handler

import (
	"fmt"
	"golang-template-api-service/app/internal/dto/common"
	"golang-template-api-service/app/internal/dto/user"
	"golang-template-api-service/app/internal/usecase"
	"golang-template-api-service/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(UserUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: UserUsecase}
}

// @Summary     Get User informations from oauth server
// @Description  Get User informations from oauth server
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "userid"
// @Success      200  {object} 	user.UserDetailResponse
// @Failure      400  {string}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v1/user/{userid}/detail [get]
func (h *UserHandler) GetUserDetails(c *fiber.Ctx) error {
	// Get the JWT token from the context
	token, ok := c.Locals("jwt_token").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "JWT token not found",
		})
	}

	fmt.Printf("Received ID from URL: %s\n", token) // Add this log line to see the value of idStr

	idStr := c.Params("id")
	fmt.Printf("Received ID from URL: %s\n", idStr) // Add this log line to see the value of idStr

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	user, err := h.UserUsecase.GetUserDetails(id, token)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	var apiResponse = common.NewResponse(user, "success")
	return c.JSON(apiResponse)
}

// @Summary     Get User informations from oauth server
// @Description  Get User informations from oauth server
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "userid"
// @Success      200  {object} 	user.UserStatusUpdate
// @Failure      400  {string}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v1/user/status [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {

	var userStatusUpdate user.UserStatusUpdate
	// Parse and populate the UpdateAdjacent struct from the request body
	if err := c.BodyParser(&userStatusUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	idStr := c.Params("id")
	fmt.Printf("Received ID from URL: %s\n", idStr) // Add this log line to see the value of idStr

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	// Get the JWT token from the context
	token, ok := c.Locals("jwt_token").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "JWT token not found",
		})
	}
	user, err := h.UserUsecase.UpdateUser(id, userStatusUpdate, token)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	var apiResponse = common.NewResponse(user, "success")
	return c.JSON(apiResponse)
}

// @Summary     Register
// @Description   Register
// @Tags         User
// @Accept       json
// @Produce      json
// @Param UserRequest body user.UserRequest true "Custom UserRequest request"
// @Success      200  {object} 	entity.User
// @Failure      400  {string}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v1/user/register [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var user user.UserRequest

	mandatoryFields := utils.HasMandatoryFields(user)

	// Parse and populate the user struct from the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// mandatory field check
	if err := utils.ValidateMandatoryFields(user, mandatoryFields); err != nil {
		return c.Status(400).JSON(map[string]interface{}{
			"error":         "mandatory field is missing or empty",
			"missing_field": err.Error(),
		})
	}

	userEntity, err := h.UserUsecase.CreateUser(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// For example, return the Register as JSON response
	var apiResponse = common.NewResponse(userEntity, "success")

	return c.JSON(apiResponse)

}

// @Summary     Get User informations from oauth server
// @Description  Get User informations from oauth server
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "userid"
// @Success      200  {object} 	user.UserDetailResponse
// @Failure      400  {string}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v1/user/{userid} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	// Get the JWT token from the context
	token, ok := c.Locals("jwt_token").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "JWT token not found",
		})
	}
	idStr := c.Params("id")
	fmt.Printf("Received ID from URL: %s\n", idStr) // Add this log line to see the value of idStr

	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	user, err := h.UserUsecase.DeleteUser(id, token)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	var apiResponse = common.NewResponse(user, "success")
	return c.JSON(apiResponse)
}
