package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;primaryKey;unique;default: gen_random_uuid();"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

// Default TableName = `users`
// TableName overrides the table name to `user`
// https://gorm.io/docs/conventions.html
func (User) TableName() string {
	return "user"
}
