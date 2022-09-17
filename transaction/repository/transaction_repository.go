package repository

import (
	"backend-core/domain"
	"context"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository struct {
	db *gorm.DB
}

func (t TransactionRepository) Store(ctx context.Context, transaction domain.Transaction) error {
	return t.db.WithContext(ctx).Create(&transaction).Error
}

func (t TransactionRepository) FetchByUserId(ctx context.Context, userId int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := t.db.WithContext(ctx).Where(domain.Asset{UserID: userId}).Preload(clause.Associations).Find(&transactions).Error
	return transactions, err
}

func (t TransactionRepository) GetTotalProfitByUserId(ctx context.Context, userId int) (float64, error) {
	var amount float64
	err := t.db.WithContext(ctx).Raw("SELECT SUM(amount) FROM transactions WHERE user_id = ? AND transaction_type_id = ?", userId, 5).Scan(&amount).Error
	return amount, err
}

func (t TransactionRepository) Fetch(ctx context.Context) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := t.db.WithContext(ctx).Preload(clause.Associations).Find(&transactions).Error
	return transactions, err
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	db.Exec("DROP TYPE IF EXISTS transaction_type;CREATE TYPE transaction_type AS ENUM ('WITHDRAW', 'DEPOSIT', 'PROFIT');")
	var result int64
	db.Table("transaction_types").Count(&result)
	if result == 0 {
		err := db.Create(domain.TransactionTypes).Error
		if err != nil {
			log.Error().Err(err)
		}
	}

	return &TransactionRepository{db}
}
