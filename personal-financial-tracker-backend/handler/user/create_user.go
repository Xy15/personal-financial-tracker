package handler_user

import (
	"personal-financial-tracker/model"
	"personal-financial-tracker/response"
	"personal-financial-tracker/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateDefaultUserCategory(ctx *fiber.Ctx, tx *gorm.DB, userID uuid.UUID) (err error) {
	var defaultCategories = []model.UserCategory{
		{ID: uuid.New(), UserID: userID, CategoryImageID: uuid.MustParse("64918708-021e-4dbb-8d12-c326e3231ba4"), Name: "Food", Type: "Expense"},
		{ID: uuid.New(), UserID: userID, CategoryImageID: uuid.MustParse("64918708-021e-4dbb-8d12-c326e3231ba4"), Name: "Fast Food", Type: "Expense"},
		{ID: uuid.New(), UserID: userID, CategoryImageID: uuid.MustParse("45528ca6-9984-446e-bb3f-1abb9eefea22"), Name: "Snack", Type: "Expense"},
		{ID: uuid.New(), UserID: userID, CategoryImageID: uuid.MustParse("ac599cc5-b565-4f79-8cd6-05ea921a10fb"), Name: "Drink", Type: "Expense"},
		{ID: uuid.New(), UserID: userID, CategoryImageID: uuid.MustParse("4f47676a-4eae-4aa2-bd77-04bcf0d6d416"), Name: "Jeans", Type: "Expense"},
		{ID: uuid.New(), UserID: userID, CategoryImageID: uuid.MustParse("d22d4b1c-4288-499d-bf5e-9c6dc2d4dbb4"), Name: "Medicine", Type: "Expense"},
		{ID: uuid.New(), UserID: userID, CategoryImageID: uuid.MustParse("64918708-021e-4dbb-8d12-c326e3231ba4"), Name: "Restaurant", Type: "Income"},
	}

	// Insert default categories for the new user
	for _, category := range defaultCategories {
		if err := tx.Create(&category).Error; err != nil {
			return err
		}
	}

	return nil
}

type CreateUserReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param body body UpdateUserReq true "User"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.Response "Bad request"
// @Failure 404 {object} response.Response "Record not found"
// @Failure 500 {object} response.Response "Internal server error"
// @Router /user [post]
func CreateUser(ctx *fiber.Ctx, db *gorm.DB) error {
	var body CreateUserReq
	if err := utils.ValidateRequestBody(&body, ctx); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	passwordHash, err := utils.HashPassword(body.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponseBody(err.Error()))
	}

	tx := db.Begin()
	userID := uuid.New()

	user := &model.User{
		ID:           userID,
		Username:     body.Username,
		Email:        body.Email,
		PasswordHash: passwordHash,
	}
	err = db.Create(user).Error
	if err != nil {
		tx.Rollback()
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	err = CreateDefaultUserCategory(ctx, tx, userID)
	if err != nil {
		tx.Rollback()
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponseBody(err.Error()))
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponseBody(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(response.DataResponseBody(user, "Successfully created user"))
}
