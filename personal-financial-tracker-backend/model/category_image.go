package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryImage struct {
	gorm.Model
	ImageID   uuid.UUID `json:"image_id" gorm:"type:uuid;primaryKey;unique;default: gen_random_uuid();"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Url string `json:"url"`
}
