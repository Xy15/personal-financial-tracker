package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID uuid.UUID `json:"transaction_id" gorm:"type:uuid;primaryKey;unique;default: gen_random_uuid();"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CategoryID      string    `json:"category_id"`
	Description     *string   `json:"description"`
	Amount          float64   `json:"amount"`
	TransactionDate time.Time `json:"transaction_date"`
	UserID          string    `json:"user_id"`

	Category Category `json:"category" gorm:"foreignKey:CategoryID;references:CategoryID"`
}

// Default TableName = `transactions`
// TableName overrides the table name to `transaction`
// https://gorm.io/docs/conventions.html
func (Transaction) TableName() string {
	return "transaction"
}
