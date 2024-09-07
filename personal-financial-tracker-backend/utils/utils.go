package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateRequestBody(body interface{}, ctx *fiber.Ctx) error {
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(body)

	return err
}
