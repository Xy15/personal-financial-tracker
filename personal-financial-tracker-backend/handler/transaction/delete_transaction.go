package handler_transaction

import (
	"personal-financial-tracker/model"
	"personal-financial-tracker/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Delete Transaction By ID
// @Description Delete Transaction By ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param transaction_id path string true "Transaction ID"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /transaction/{transaction_id} [delete]
func DeleteTransactionByID(ctx *fiber.Ctx, db *gorm.DB) error {
	transactionID := ctx.Params("transaction_id")

	err := db.Delete(&model.Transaction{}, transactionID).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponseBody("Successfully deleted transaction"))
}
