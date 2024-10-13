package handler_user_category

import (
	"personal-financial-tracker/response"
	"personal-financial-tracker/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Delete User Category By ID
// @Description Delete User Category By ID
// @Tags User Category
// @Accept json
// @Produce json
// @Param user_category_id path string true "User Category ID"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/category/{user_category_id} [delete]
func DeleteUserCategoryByID(ctx *fiber.Ctx, db *gorm.DB) error {
	userCategoryID := ctx.Params("user_category_id")

	err := service.DeleteUserCategoryByID(userCategoryID, db)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponseBody("Successfully Deleted User Category"))
}
