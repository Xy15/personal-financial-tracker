package router

import (
	handler_auth "personal-financial-tracker/handler/auth"
	handler_transaction "personal-financial-tracker/handler/transaction"
	handler_user "personal-financial-tracker/handler/user"
	handler_user_category "personal-financial-tracker/handler/user_category"
	"personal-financial-tracker/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(app *fiber.App, db *gorm.DB) {
	user := app.Group("/user")

	user.Get("/:user_id", func(c *fiber.Ctx) error {
		return handler_user.GetUserByID(c, db)
	})
	user.Post("/", func(c *fiber.Ctx) error {
		return handler_user.CreateUser(c, db)
	})
	user.Patch("/:user_id", func(c *fiber.Ctx) error {
		return handler_user.UpdateUserByID(c, db)
	})
	user.Delete("/:user_id", func(c *fiber.Ctx) error {
		return handler_user.DeleteUserByID(c, db)
	})
}

func UserCategoryRoutes(app *fiber.App, db *gorm.DB) {
	user := app.Group("/user")

	user.Get("/:user_id/category", func(c *fiber.Ctx) error {
		return handler_user_category.GetUserCategoriesByUserID(c, db)
	})
	user.Post("/:user_id/category", func(c *fiber.Ctx) error {
		return handler_user_category.CreateUserCategory(c, db)
	})
	user.Patch("/category/:user_category_id", func(c *fiber.Ctx) error {
		return handler_user_category.UpdateUserCategoryByID(c, db)
	})
	user.Delete("/category/:user_category_id", func(c *fiber.Ctx) error {
		return handler_user_category.DeleteUserCategoryByID(c, db)
	})
}

func TransactionRoutes(app *fiber.App, db *gorm.DB) {
	transaction := app.Group("/transaction")

	transaction.Get("/today", middleware.ValidateBearerToken(), func(c *fiber.Ctx) error {
		return handler_transaction.GetTodayTransactions(c, db)
	})
	transaction.Get("/user/:user_id", func(c *fiber.Ctx) error {
		return handler_transaction.GetTransactionsByUserID(c, db)
	})
	transaction.Post("/", func(c *fiber.Ctx) error {
		return handler_transaction.CreateTransaction(c, db)
	})
	transaction.Patch("/:transaction_id", func(c *fiber.Ctx) error {
		return handler_transaction.UpdateTransactionByID(c, db)
	})
	transaction.Delete("/:transaction_id", func(c *fiber.Ctx) error {
		return handler_transaction.DeleteTransactionByID(c, db)
	})
}

func TokenRoutes(app *fiber.App, db *gorm.DB) {
	token := app.Group("/token")

	token.Get("/refresh", func(c *fiber.Ctx) error {
		return handler_auth.RefreshToken(c, db)
	})
}
