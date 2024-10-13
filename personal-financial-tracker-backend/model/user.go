package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;unique;default: gen_random_uuid();"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`

	Transactions []Transaction `json:"transactions" gorm:"foreignKey:UserID;references:ID"`
}
