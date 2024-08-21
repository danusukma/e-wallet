package controllers

import (
	"e-wallet/database"
	"e-wallet/models"
	"e-wallet/utils"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func CustomerTransfer(c *fiber.Ctx) error {

	//Get username from token
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

	var CustBalanceFROM models.ResponseCustomerBalance
	err = database.DB.Raw("SELECT Id, UserName, Balance FROM customers WHERE Username = ? LIMIT 1", userName).Scan(&CustBalanceFROM).Error

	//Request body
	customerTF := new(models.RequestCustomerTransfer)

	if err := c.BodyParser(customerTF); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(customerTF)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	//Checking source balance
	if CustBalanceFROM.Balance < customerTF.Amount {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Insufficient balance",
		})
	}

	var CustBalanceTo models.ResponseCustomerBalance
	err = database.DB.Raw("SELECT Id, UserName, Balance FROM customers WHERE Username = ? LIMIT 1", customerTF.ToUserName).Scan(&CustBalanceTo).Error

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    404,
			"message": "Destination user not found",
		})
	}

	//Checking destination username
	if CustBalanceTo.UserName == "" {
		return c.Status(400).JSON(fiber.Map{
			"code":    404,
			"message": "Destination user not found",
		})
	}

	//TypeTransaction 0=top up 1=transfer
	newWallet := models.WalletTransaction{
		TypeTransaction: 1,
		FromId:          CustBalanceFROM.Id,
		ToId:            CustBalanceTo.Id,
		Amount:          customerTF.Amount,
	}

	//Insert to wallet transaction
	database.DB.Create(&newWallet)

	//Update balance source
	database.DB.Debug().Model(&models.Customer{}).Where("id = ?", CustBalanceFROM.Id).Updates(map[string]interface{}{
		"balance": CustBalanceFROM.Balance - customerTF.Amount,
	})

	//Update balance destination
	database.DB.Debug().Model(&models.Customer{}).Where("id = ?", CustBalanceTo.Id).Updates(map[string]interface{}{
		"balance": CustBalanceTo.Balance + customerTF.Amount,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    204,
		"message": "Transfer success",
	})
}

func CustomerBalanceTopUp(c *fiber.Ctx) error {

	//Get username from token
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

	var CustBalanceFROM models.ResponseCustomerBalance
	err = database.DB.Raw("SELECT Id, UserName, Balance FROM customers WHERE Username = ? LIMIT 1", userName).Scan(&CustBalanceFROM).Error

	//Request body
	customerTopUp := new(models.RequestCustomerBalanceTopUp)

	if err := c.BodyParser(customerTopUp); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(customerTopUp)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid topup amount",
			"error":   errValidate.Error(),
		})
	}

	//TypeTransaction 0=top up 1=transfer
	newWallet := models.WalletTransaction{
		TypeTransaction: 0,
		FromId:          CustBalanceFROM.Id,
		ToId:            CustBalanceFROM.Id,
		Amount:          customerTopUp.Amount,
	}

	//Insert to wallet transaction
	database.DB.Create(&newWallet)

	//Update balance source
	database.DB.Debug().Model(&models.Customer{}).Where("id = ?", CustBalanceFROM.Id).Updates(map[string]interface{}{
		"balance": CustBalanceFROM.Balance + customerTopUp.Amount,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    204,
		"message": "Topup success",
	})
}

func GetTopTransferTransaction(c *fiber.Ctx) error {
	//Get username from token
	token := c.Get("Authorization") // Get token header
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	_, err := utils.DecodeToken(token)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var CustTopTransaction []models.ResponseCustomerTopTransaction
	err = database.DB.Raw("SELECT a.UserName, SUM(b.Amount) AS TransactedValue FROM customers a INNER JOIN wallet_transactions b ON a.id = b.FromId WHERE b.TypeTransaction  = 1 GROUP BY a.UserName  ORDER BY TransactedValue DESC LIMIT 10").Scan(&CustTopTransaction).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": CustTopTransaction,
	})

}

func GetTopTransaction(c *fiber.Ctx) error {
	//Get username from token
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

	var CustTopTransaction []models.ResponseCustomerTopTransaction
	err = database.DB.Raw("SELECT a.UserName, MAX(b.Amount) AS TransactedValue FROM customers a INNER JOIN wallet_transactions b ON a.id = b.FromId WHERE b.TypeTransaction  = 1 AND a.UserName = ? GROUP BY a.UserName  ORDER BY TransactedValue DESC LIMIT 10", userName).Scan(&CustTopTransaction).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"message": CustTopTransaction,
	})

}
