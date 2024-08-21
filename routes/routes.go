package routes

import (
	"e-wallet/controllers"
	"e-wallet/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	//Get Method
	r.Get("/", func(c *fiber.Ctx) error { return c.SendString("e-Wallet api health check") })
	r.Get("/balance_read", middleware.UserAuth, controllers.GetCustomerBalance)
	r.Get("/top_users", middleware.UserAuth, controllers.GetTopTransferTransaction)
	r.Get("/top_transaction_per_user", middleware.UserAuth, controllers.GetTopTransaction)

	//Post Method
	r.Post("/create_user", controllers.CreateCustomer)
	r.Post("/login", controllers.Login)
	r.Post("/transfer", middleware.UserAuth, controllers.CustomerTransfer)
	r.Post("/balance_topup", middleware.UserAuth, controllers.CustomerBalanceTopUp)
}
