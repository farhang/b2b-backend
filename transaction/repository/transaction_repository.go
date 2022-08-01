package repository

import (
	"backend-core/domain"
	"context"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func (t TransactionRepository) Store(ctx context.Context, transaction domain.Transaction) error {
	return t.db.WithContext(ctx).Create(&transaction).Error
}

func (t TransactionRepository) FetchByUserId(ctx context.Context, userId int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := t.db.WithContext(ctx).Where(domain.Asset{UserID: userId}).Find(&transactions).Error
	return transactions, err
}
func (t TransactionRepository) Fetch(ctx context.Context) ([]domain.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	db.Exec("CREATE TYPE transaction_type AS ENUM (?,?,?)", domain.WITHDRAW, domain.DEPOSIT, domain.PROFIT)
	return &TransactionRepository{db}
}
