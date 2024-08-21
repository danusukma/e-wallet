package controllers

import (
	"e-wallet/database"
	"e-wallet/models"
	"e-wallet/utils"
	"time"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(models.Login)

	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to Login",
			"error":   errValidate.Error(),
		})
	}

	var customer models.Customer

	// Check Validation Email
	errUserName := database.DB.Debug().First(&customer, "UserName = ?", loginRequest.UserName).Error
	if errUserName != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Check Password Validation
	isValid := utils.CheckHash(loginRequest.Password, customer.Password)
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Password is not valid",
		})
	}

	// Generate JWT
	claims := jwt.MapClaims{}
	// Claims or Payloads Definitions
	claims["name"] = customer.FullName
	claims["username"] = customer.UserName
	claims["exp"] = time.Now().Add(2 * time.Minute).Unix()

	token, errGenerate := utils.GenerateToken(&claims)

	if errGenerate != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Token is not valid",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
