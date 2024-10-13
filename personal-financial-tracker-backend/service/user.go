package service

import (
	"personal-financial-tracker/model"
	"personal-financial-tracker/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserByID(userID string, db *gorm.DB) (*model.User, error) {
	var user *model.User
	err := db.Where("id = ?", userID).
		First(&user).Error

	return user, err
}

type CreateUserReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserByIDReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(body *CreateUserReq, db *gorm.DB) (*model.User, error) {
	passwordHash, err := utils.HashPassword(body.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:           uuid.New(),
		Username:     body.Username,
		Email:        body.Email,
		PasswordHash: passwordHash,
	}

	err = db.Create(user).Error

	return user, err
}

func UpdateUserByID(userID string, body *UpdateUserByIDReq, db *gorm.DB) error {
	passwordHash, err := utils.HashPassword(body.Password)
	if err != nil {
		return err
	}

	user := &model.User{
		Username:     body.Username,
		Email:        body.Email,
		PasswordHash: passwordHash,
	}
	err = db.Model(&model.User{}).Where("id = ?", userID).Updates(user).Error

	return err
}

func DeleteUserByID(userID string, db *gorm.DB) error {
	err := db.Delete(&model.User{}, userID).Error

	return err
}
