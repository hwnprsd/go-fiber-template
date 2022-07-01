package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

// Assign the request body to the given interface and then check for validity
func ValidateBody(data interface{}, c *fiber.Ctx) (bool, []*ErrorResponse) {
	if err := c.BodyParser(&data); err != nil {
		errData := make([]*ErrorResponse, 1)
		error := ErrorResponse{
			FailedField: "nil",
			Tag:         "Parse Error",
			Value:       "nil",
		}
		errData = append(errData, &error)
		return true, errData
	}

	var validate = validator.New()
	var errors []*ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	if len(errors) > 0 {
		return true, errors
	}
	return false, nil
}

// Throw a validation error
func Throw(c *fiber.Ctx, errors interface{}) error {
	return c.Status(fiber.ErrBadRequest.Code).JSON(errors)
}

// return just an object id
func ObjId(s string) primitive.ObjectID {
	o, _ := primitive.ObjectIDFromHex(s)
	return o
}
