package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SendSuccess(c *fiber.Ctx, status int, data interface{}, message ...string) error {
	response := fiber.Map{
		"success": true,
		"data":    data,
	}

	if len(message) > 0 {
		response["message"] = message[0]
	}

	return c.Status(status).JSON(response)
}

func SendError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"error":   message,
	})
}

func SendValidationError(c *fiber.Ctx, err error) error {
	errors := FormatValidationError(err)
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"errors":  errors,
	})
}

func FormatValidationError(err error) []string {
	var validationErrors validator.ValidationErrors
	errorMessages := make([]string, 0)

	if errors.As(err, &validationErrors) {
		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("%s is required", e.Field()))
			case "email":
				errorMessages = append(errorMessages, "Invalid email format")
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param()))
			case "max":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must not exceed %s characters", e.Field(), e.Param()))
			case "eqfield":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be equal to %s", e.Field(), e.Param()))
			case "password_validator":
				errorMessages = append(errorMessages, "Password must contain at least one uppercase letter, one number, and one special character")
			}
		}
	}

	return errorMessages
}
