package users

import (
	"flaq.club/backend/utils"
	"github.com/gofiber/fiber/v2"
)

func UserController(app *fiber.App) {
	userRouter := app.Group("/users")
	setupRoutes(userRouter)
}

func setupRoutes(router fiber.Router) {
	router.Get("/all", func(c *fiber.Ctx) error {
		return utils.GetRequestHandler(c, GetAllUsers, utils.RequestBody{})
	})

	router.Post("/one", func(c *fiber.Ctx) error {
		body := new(OneUserDto)
		return utils.PostRequestHandler(c, GetOneUser, utils.RequestBody{Data: body})
	})
}
