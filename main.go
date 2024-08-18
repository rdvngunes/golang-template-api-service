package main

import (
	"golang-template-api-service/app/config"
	"golang-template-api-service/app/router"
	storage "golang-template-api-service/app/storage"
	"golang-template-api-service/app/utils"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

var configuration *config.Configuration

func init() {
	configuration = config.LoadViperConfig()
}
func main() {

	utils.SetupLogs() //logging layer

	app := fiber.New(fiber.Config{
		ReadBufferSize: 4096 * 10,
	})

	db, err := storage.SetupDatabase() // database layer
	if err != nil {
		errMsg := fmt.Sprintf("Error setting up database: %v", err)
		utils.ErrorLog(errMsg)
		// log.Fatalf("Error setting up database: %v", err)
	}
	//utils.InitCacheRedis() // caching layer

	router.SetupMainRouter(app, db, utils.NewHTTPClient()) // router layer

	app.Listen(fmt.Sprintf(":%d", configuration.App.Port))
}
