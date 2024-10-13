package service

import (
	"personal-financial-tracker/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetTransactionByID(transactionID string, db *gorm.DB) (*model.Transaction, error) {
	var transaction *model.Transaction
	err := db.Where("id = ?", transactionID).
		First(&transaction).Error

	return transaction, err
}

func GetTransactionsByUserID(userID string, db *gorm.DB) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := db.Where("user_id = ?", userID).
		Order("transaction_date DESC").
		Find(&transactions).Error

	return transactions, err
}

func GetTransactionsByUserIDGroupedByDate(userID string, db *gorm.DB) (map[string][]*model.Transaction, error) {
	var transactions []*model.Transaction
	err := db.Where("user_id = ?", userID).
		Order("transaction_date DESC").
		Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	// Group transactions by transaction_date
	groupedTransactions := make(map[string][]*model.Transaction)
	for _, transaction := range transactions {
		// Format date as "DD-MM-YYYY"
		transactionDate := transaction.TransactionDate
		dateStr := transactionDate.Format("02/01/2006")
		groupedTransactions[dateStr] = append(groupedTransactions[dateStr], transaction)
	}
	return groupedTransactions, err
}

type CreateTransactionReq struct {
	UserCategoryID  string    `json:"user_category_id" validate:"required"`
	Description     *string   `json:"description"`
	Amount          float64   `json:"amount" validate:"required"`
	TransactionDate time.Time `json:"transaction_date" validate:"required"`
	UserID          string    `json:"user_id" validate:"required"`
}

func CreateTransaction(body *CreateTransactionReq, db *gorm.DB) (*model.Transaction, error) {
	userCategory, err := GetUserCategoryByID(body.UserCategoryID, db)
	if err != nil {
		return nil, err
	}

	transaction := &model.Transaction{
		ID:               uuid.New(),
		CategoryImageUrl: userCategory.CategoryImage.ImageUrl,
		CategoryName:     userCategory.Name,
		Description:      body.Description,
		Amount:           body.Amount,
		TransactionDate:  body.TransactionDate,
		UserID:           body.UserID,
	}

	err = db.Create(transaction).Error

	return transaction, err
}
