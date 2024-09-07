package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	CategoryID uuid.UUID `json:"category_id" gorm:"type:uuid;primaryKey;unique;default: gen_random_uuid();"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	UserID        *string       `json:"user_id"`
	ImageID       string        `json:"image_id"`
	Name          string        `json:"name"`
	Type          string        `json:"type"`
	CategoryImage CategoryImage `json:"category_image" gorm:"foreignKey:ImageID;references:ImageID"`
}

func (Category) TableName() string {
	return "category"
}
