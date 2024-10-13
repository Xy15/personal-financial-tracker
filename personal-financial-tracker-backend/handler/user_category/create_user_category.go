package handler_user_category

import (
	"personal-financial-tracker/response"
	"personal-financial-tracker/service"
	"personal-financial-tracker/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Create User Category
// @Description Create User Category
// @Tags User Category
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param body body service.CreateUserCategoryReq true "User Category"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/{user_id}/category [post]
func CreateUserCategory(ctx *fiber.Ctx, db *gorm.DB) error {
	userID := ctx.Params("user_id")

	var body service.CreateUserCategoryReq
	if err := utils.ValidateRequestBody(&body, ctx); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	userCategory, err := service.CreateUserCategory(userID, &body, db)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponseBody(userCategory, "Successfully Created User Category"))
}
