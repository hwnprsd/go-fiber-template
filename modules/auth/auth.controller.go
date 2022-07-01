package auth

import (
	. "flaq.club/backend/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthController(app *fiber.App) {
	authRouter := app.Group("/auth")
	setupRoutes(authRouter)
}

func setupRoutes(router fiber.Router) {

	router.Post("/login", func(c *fiber.Ctx) error {
		body := new(LoginDto)
		return PostRequestHandler(c, Login, RequestBody{Data: body})
	})

	router.Post("/signup", func(c *fiber.Ctx) error {
		body := new(SignupDto)
		return PostRequestHandler(c, Signup, RequestBody{Data: body})
	})
}
