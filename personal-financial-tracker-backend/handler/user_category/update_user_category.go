package handler_user_category

import (
	"personal-financial-tracker/response"
	"personal-financial-tracker/service"
	"personal-financial-tracker/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Update User Category
// @Description Update User Category
// @Tags User Category
// @Accept json
// @Produce json
// @Param user_category_id path string true "User Category ID"
// @Param body body service.UpdateUserCategoryByIDReq true "User Category"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/category/{user_category_id} [patch]
func UpdateUserCategoryByID(ctx *fiber.Ctx, db *gorm.DB) error {
	userCategoryID := ctx.Params("user_category_id")

	var body *service.UpdateUserCategoryByIDReq
	if err := utils.ValidateRequestBody(body, ctx); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	err := service.UpdateUserCategoryByID(userCategoryID, body, db)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponseBody("Successfully Updated User Category"))
}
