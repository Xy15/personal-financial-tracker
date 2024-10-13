package handler_transaction

import (
	"personal-financial-tracker/model"
	"personal-financial-tracker/response"
	"personal-financial-tracker/service"
	"personal-financial-tracker/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UpdateTransactionReq struct {
	// CategoryImageUrl string    `json:"category_image_url"`
	// CategoryName     string    `json:"category_name"`
	UserCategoryID  string    `json:"user_category_id"`
	Description     *string   `json:"description"`
	Amount          float64   `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
	UserID          string    `json:"user_id"`
}

// @Summary Update Transaction By ID
// @Description Update Transaction By ID
// @Tags Transaction
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param transaction_id path string true "Transaction ID"
// @Param body body UpdateTransactionReq true "Transaction"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /transaction/{transaction_id} [patch]
func UpdateTransactionByID(ctx *fiber.Ctx, db *gorm.DB) error {
	transactionID := ctx.Params("transaction_id")
	var body UpdateTransactionReq
	if err := utils.ValidateRequestBody(&body, ctx); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	userCategory, err := service.GetUserCategoryByID(body.UserCategoryID, db)

	transaction := &model.Transaction{
		CategoryImageUrl: userCategory.CategoryImage.ImageUrl,
		CategoryName:     userCategory.Name,
		Description:      body.Description,
		Amount:           body.Amount,
		TransactionDate:  body.TransactionDate,
	}

	err = db.Model(&model.Transaction{}).Where("id = ?", transactionID).Updates(transaction).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponseBody(transaction, "Successfully updated transaction"))
}
