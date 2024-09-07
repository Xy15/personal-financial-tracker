package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionID uuid.UUID `json:"transaction_id" gorm:"type:uuid;primaryKey;unique;default: gen_random_uuid();"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CategoryID  string  `json:"category_id"`
	Description *string `json:"description"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	UserID      string  `json:"user_id"`

	Category
}
