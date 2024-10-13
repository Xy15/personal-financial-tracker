package handler_auth

import (
	"personal-financial-tracker/response"
	"personal-financial-tracker/service"
	"personal-financial-tracker/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary Refresh Token
// @Description Refresh Token
// @Tags Authentication
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param body body service.RefreshTokenReq true "Refresh Token"
// @Success 200 {object} response.Response{data=service.RefreshTokenRes} "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /token/refresh [post]
func RefreshToken(ctx *fiber.Ctx, db *gorm.DB) error {
	var body service.RefreshTokenReq
	if err := utils.ValidateRequestBody(&body, ctx); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	refreshTokenClaims, err := service.ParseRefreshToken(body.RefreshToken)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	user, err := service.GetUserByID(refreshTokenClaims.UserID, db)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	newAccessToken, err := service.GenerateAccessToken(user.ID.String(), user.Email, user.Username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponseBody(service.RefreshTokenRes{NewAccessToken: *newAccessToken}, "Successfully generated new access token"))
}
