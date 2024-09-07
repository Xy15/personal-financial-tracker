package handler

import (
	"personal-financial-tracker/model"
	"personal-financial-tracker/response"
	"personal-financial-tracker/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetTransactionsByUserID(ctx *fiber.Ctx, gdb *gorm.DB) error {
	userID := ctx.Params("user_id")

	var transactions []model.Transaction
	err := gdb.Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponseBody(transactions, "Successfully get transactions"))
}

type IncomingCreateTransaction struct {
	CategoryID  string  `json:"category_id" validate:"required"`
	Description *string `json:"description" validate:"-"`
	Type        string  `json:"type" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	UserID      string  `json:"user_id" validate:"required"`
}

func CreateTransaction(ctx *fiber.Ctx, gdb *gorm.DB) error {
	var body IncomingCreateTransaction
	if err := utils.ValidateRequestBody(&body, ctx); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	// var transaction []model.Transaction
	// save certificate media
	transaction := &model.Transaction{
		TransactionID: uuid.New(),
		CategoryID:    body.CategoryID,
		Description:   body.Description,
		Type:          body.Type,
		Amount:        body.Amount,
		UserID:        body.UserID,
	}

	err := gdb.Create(transaction).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponseBody(transaction, "Successfully created transaction"))
}
