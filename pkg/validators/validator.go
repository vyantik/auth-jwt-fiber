package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	validate := validator.New()
	validate.RegisterValidation("strongPassword", ValidateStrongPassword)

	return &Validator{
		validate: validate,
	}
}

func (v *Validator) ValidateRequest(c *fiber.Ctx, req any) error {
	if err := v.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
			"errors":  err.Error(),
		})
	}

	return nil
}
