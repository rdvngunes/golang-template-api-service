package router

import (
	"golang-template-api-service/app/internal/handler"
	"golang-template-api-service/app/internal/repository"
	"golang-template-api-service/app/internal/usecase"
	"golang-template-api-service/app/middleware"
	"golang-template-api-service/app/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupUserRoutes(user_router fiber.Router, db *gorm.DB, httpClient *utils.HTTPClient) {

	user_repository := repository.NewUserRepository(db)
	user_usecase := usecase.NewUserUsecase(user_repository, httpClient)
	user_handler := handler.NewUserHandler(user_usecase)

	user := user_router.Group("/user")

	user.Get("/:Id/detail", middleware.AuthMiddlewareWithScopes("first-scope", "second-scope", "third-scope", "fourth-scope"), user_handler.GetUserDetails)
	user.Delete("/:Id", middleware.AuthMiddlewareWithScopes("third-scope"), user_handler.DeleteUser)
	user.Post("/register", user_handler.Register)
	user.Put("/:Id", middleware.AuthMiddlewareWithScopes("third-scope"), user_handler.UpdateUser)
}
