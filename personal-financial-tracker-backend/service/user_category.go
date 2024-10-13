package service

import (
	"personal-financial-tracker/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserCategoryByID(userCategoryID string, db *gorm.DB) (*model.UserCategory, error) {
	var userCategory *model.UserCategory
	err := db.Where("id = ?", userCategoryID).
		Preload("CategoryImage").
		First(&userCategory).Error

	return userCategory, err
}

type CreateUserCategoryReq struct {
	CategoryImageID string `json:"category_image_id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Type            string `json:"type" validate:"required"`
}

type UpdateUserCategoryByIDReq struct {
	CategoryImageID string `json:"category_image_id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
}

func CreateUserCategory(userID string, body *CreateUserCategoryReq, db *gorm.DB) (*model.UserCategory, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	categoryImageUUID, err := uuid.Parse(body.CategoryImageID)
	if err != nil {
		return nil, err
	}

	userCategory := &model.UserCategory{
		ID:              uuid.New(),
		UserID:          userUUID,
		CategoryImageID: categoryImageUUID,
		Name:            body.Name,
		Type:            body.Type,
	}

	err = db.Create(userCategory).Error
	if err != nil {
		return nil, err
	}

	return userCategory, nil
}

func UpdateUserCategoryByID(userCategoryID string, body *UpdateUserCategoryByIDReq, db *gorm.DB) error {
	categoryImageUUID, err := uuid.Parse(body.CategoryImageID)
	if err != nil {
		return err
	}

	userCategory := &model.UserCategory{
		CategoryImageID: categoryImageUUID,
		Name:            body.Name,
		Type:            body.Type,
	}
	err = db.Model(&model.UserCategory{}).Where("id = ?", userCategoryID).Updates(userCategory).Error

	return err
}

func DeleteUserCategoryByID(userCategoryID string, db *gorm.DB) error {
	err := db.Delete(&model.UserCategory{}, userCategoryID).Error

	return err
}
