package router

import (
	"personal-financial-tracker/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TransactionRoutes(app *fiber.App, db *gorm.DB) {
	transaction := app.Group("/transaction")

	transaction.Get("/user/:user_id", func(c *fiber.Ctx) error {
		return handler.GetTransactionsByUserID(c, db)
	})
	transaction.Post("/", func(c *fiber.Ctx) error {
		return handler.CreateTransaction(c, db)
	})
}
