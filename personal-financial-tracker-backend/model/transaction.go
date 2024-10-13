package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;unique;default: gen_random_uuid();"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CategoryImageUrl string    `json:"category_image_url"`
	CategoryName     string    `json:"category_name"`
	Description      *string   `json:"description"`
	Amount           float64   `json:"amount"`
	TransactionDate  time.Time `json:"transaction_date"`
	UserID           string    `json:"user_id"`
}
