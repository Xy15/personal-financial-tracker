package handler_transaction

import (
	"personal-financial-tracker/response"
	"personal-financial-tracker/service"
	"personal-financial-tracker/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Create Transaction
// @Description Create Transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param body body service.CreateTransactionReq true "Transaction"
// @Success 200 {object} response.Response{data=model.Transaction} "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /transaction [post]
func CreateTransaction(ctx *fiber.Ctx, db *gorm.DB) error {
	var body service.CreateTransactionReq
	if err := utils.ValidateRequestBody(&body, ctx); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	transaction, err := service.CreateTransaction(&body, db)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponseBody(transaction, "Successfully created transaction"))
}
