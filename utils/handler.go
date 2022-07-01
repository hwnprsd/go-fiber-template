package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	Data interface{}
	User interface{}
}

type PostHandler func(data RequestBody) interface{}
type GetHandler func(data RequestBody) interface{}

func handlePanic(c *fiber.Ctx) {
	// Panic handler
	if err := recover(); err != nil {
		fmt.Println("We survived a panic")
		fmt.Println(err)
		c.Status(500).JSON(fiber.Map{
			"statusCode": 500,
			"message":    "A server panic has occured",
			"error":      err,
		})
	}
}

func PostRequestHandler(c *fiber.Ctx, actualHandler PostHandler, body RequestBody) error {
	// Handle Panics
	defer handlePanic(c)

	if b, err := ValidateBody(body.Data, c); b {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"statusCode": 500,
			"error":      err,
		})
	}
	res := actualHandler(body)
	return c.Status(200).JSON(fiber.Map{
		"statusCode": 200,
		"data":       res,
	})
}

func GetRequestHandler(c *fiber.Ctx, actualHandler GetHandler, body RequestBody) error {
	// Handle Panics
	defer handlePanic(c)
	res := actualHandler(body)
	return c.Status(200).JSON(fiber.Map{
		"statusCode": 200,
		"data":       res,
	})
}
