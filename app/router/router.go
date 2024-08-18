package router

import (
	"golang-template-api-service/app/internal/router"
	"golang-template-api-service/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func SetupMainRouter(app *fiber.App, db *gorm.DB, httpClient *utils.HTTPClient) {

	// @Summary     Health
	// @Description  Health
	// @Tags         User
	// @Accept       json
	// @Produce      json
	// @Success      200  {string} 	string
	// @Failure      400  {string}  error
	// @Failure      404  {object}  error
	// @Failure      500  {object}  error
	// @Router       /api/v1/health [get]
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "success", "message": "Healthy", "data": "Ok"})
	})

	api := app.Group("/api/v1/")
	app.Static("/swagger.yaml", "./docs/swagger.yaml")

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger.yaml",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "list",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "[[Client-Id]]",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
	}))

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(redirect.New(redirect.Config{
		Rules:      map[string]string{"/": "/swagger/index.html"},
		StatusCode: 301,
	}))

	router.SetupUserRoutes(api, db, httpClient)
}
