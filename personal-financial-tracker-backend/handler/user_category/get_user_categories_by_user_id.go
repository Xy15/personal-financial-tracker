package handler_user_category

import (
	"personal-financial-tracker/model"
	"personal-financial-tracker/response"
	"personal-financial-tracker/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Get User Categories By User ID
// @Description Get User Categories By User ID
// @Tags User Category
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]model.UserCategory} "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/{user_id}/category [get]
func GetUserCategoriesByUserID(ctx *fiber.Ctx, db *gorm.DB) error {
	userID := ctx.Params("user_id")

	var userCategories []model.UserCategory
	err := db.Where("user_id = ?", userID).
		Preload("CategoryImage").
		Order("name ASC").
		Find(&userCategories).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponseBody(userCategories, "Successfully get User Categories"))
}

// @Summary Get Transactions By User ID
// @Description Get Transaction By User ID \n grouped by date
//
// @Tags Transaction
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} response.Response{data=map[string][]model.Transaction} "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /transaction/user/{user_id} [get]
func GetTransactionsByUserID(ctx *fiber.Ctx, db *gorm.DB) error {
	userID := ctx.Params("user_id")

	groupedTransactions, err := service.GetTransactionsByUserIDGroupedByDate(userID, db)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponseBody(groupedTransactions, "Successfully get transactions"))
}
