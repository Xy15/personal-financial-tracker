package handler_user

import (
	"personal-financial-tracker/model"
	"personal-financial-tracker/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Delete User By ID
// @Description Delete User By ID
// @Tags User
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param user_id path string true "User ID"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/{user_id} [delete]
func DeleteUserByID(ctx *fiber.Ctx, db *gorm.DB) error {
	userID := ctx.Params("user_id")

	err := db.Delete(&model.User{}, userID).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponseBody("Successfully deleted user"))
}
