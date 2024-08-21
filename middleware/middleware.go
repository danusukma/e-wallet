package middleware

import (
	"e-wallet/utils"

	"github.com/gofiber/fiber/v2"
)

func UserAuth(c *fiber.Ctx) error {
	token := c.Get("Authorization") // Get token header
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims, err := utils.DecodeToken(token)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	username := claims["username"].(string)
	if username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "You are not Unauthorized",
		})
	}

	return c.Next()
}
