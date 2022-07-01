package main

import (
	"flaq.club/backend/configs"
	_ "flaq.club/backend/docs"
	"flaq.club/backend/modules/auth"
	"flaq.club/backend/modules/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title           Flaq Club API
// @version         1.0
// @description     Flaq Club lives here
// @termsOfService  http://swagger.io/terms/
// @contact.name    API Support
// @contact.email   fiber@swagger.io
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8080
// @BasePath        /
func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.SendString("POOPED")
		},
	})

	configs.ConnectDB()

	auth.AuthController(app)
	users.UserController(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Listen(":8080")

}

func setupRoutes() {
}
