package middleware

import (
	"personal-financial-tracker/consts"
	"personal-financial-tracker/response"
	"personal-financial-tracker/service"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ValidateBearerToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")

		var token string
		if strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			errMsg := "Unauthorizated"
			return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponseBody(errMsg))
		}

		accessClaims, err := service.ParseAccessToken(token)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponseBody(err.Error()))
		}

		// set ctx var
		ctx.Locals(consts.EMAIL, accessClaims.Email)
		ctx.Set(consts.EMAIL, accessClaims.Email)
		ctx.Set(consts.USER_ID, accessClaims.UserID)
		ctx.Set(consts.USERNAME, accessClaims.Username)

		return ctx.Next()
	}

}
