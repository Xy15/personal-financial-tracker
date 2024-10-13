package model

import (
	"time"

	"github.com/google/uuid"
)

type UserCategory struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;unique;default: gen_random_uuid();"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	UserID          uuid.UUID     `json:"user_id"`
	CategoryImageID uuid.UUID     `json:"category_image_id"`
	Name            string        `json:"name"`
	Type            string        `json:"type"`
	CategoryImage   CategoryImage `json:"category_image" gorm:"foreignKey:ID;references:CategoryImageID"`
}
