package controllers

import (
	"e-wallet/database"
	"e-wallet/models"
	"e-wallet/utils"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func CreateCustomer(c *fiber.Ctx) error {
	customer := new(models.Customer)
	var countRow int64

	if err := c.BodyParser(customer); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(customer)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newCustomer := models.Customer{
		UserName: customer.UserName,
		FullName: customer.FullName,
	}

	countRow = 0
	database.DB.Raw("SELECT id, username FROM customers WHERE username = ?", customer.UserName).Count(&countRow)

	if countRow > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    409,
			"message": "username already exist",
		})
	}

	hashPassword, err := utils.HashPassword(customer.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "status internal server error",
		})
	}

	newCustomer.Password = hashPassword

	database.DB.Create(&newCustomer)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": "New user has been created",
	})
}

func GetCustomerBalance(c *fiber.Ctx) error {

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

	userName := claims["username"].(string)

	var CustBalance models.ResponseCustomerBalance
	err = database.DB.Raw("SELECT Id, UserName, Balance FROM customers WHERE Username = ? LIMIT 1", userName).Scan(&CustBalance).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    500,
			"message": "status internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"balance": CustBalance.Balance,
		"message": CustBalance,
	})

}
