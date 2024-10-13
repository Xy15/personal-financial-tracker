package model

import (
	"time"

	"github.com/google/uuid"
)

type CategoryImage struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;unique;default: gen_random_uuid();"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CategoryID uuid.UUID `json:"category_id"`
	ImageUrl   string    `json:"image_url"`
}
