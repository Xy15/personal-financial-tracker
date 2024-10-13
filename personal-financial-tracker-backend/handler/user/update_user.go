package handler_user

import (
	"personal-financial-tracker/model"
	"personal-financial-tracker/response"
	"personal-financial-tracker/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UpdateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary Update User By ID
// @Description Update User By ID
// @Tags User
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param user_id path string true "User ID"
// @Param body body UpdateUserReq true "User"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user/{user_id} [patch]
func UpdateUserByID(ctx *fiber.Ctx, db *gorm.DB) error {
	userID := ctx.Params("user_id")

	var body UpdateUserReq
	err := utils.ValidateRequestBody(&body, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	var passwordHash string
	if body.Password != "" {
		passwordHash, err = utils.HashPassword(body.Password)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponseBody(err.Error()))
		}
	}

	user := &model.User{
		Username:     body.Username,
		Email:        body.Email,
		PasswordHash: passwordHash,
	}
	err = db.Model(&model.User{}).Where("id = ?", userID).Updates(user).Error
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponseBody(user, "Successfully updated user"))
}
