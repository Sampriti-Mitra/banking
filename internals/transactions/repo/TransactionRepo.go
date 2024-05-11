package repo

import (
	"github.com/Sampriti-Mitra/transactions/internals/transactions/models"
	"gorm.io/gorm"
)

type ITransactionRepo interface {
	CreateTransaction(transaction *models.Transaction) error
	CreateAccount(account *models.Account) error
	FetchAccount(accountId int64) (*models.Account, error)
}

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) (*TransactionRepo, error) {
	err := db.AutoMigrate(&models.Account{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.Transaction{})
	if err != nil {
		return nil, err
	}
	return &TransactionRepo{db}, nil
}

func (tr *TransactionRepo) CreateTransaction(transaction *models.Transaction) error {
	return tr.db.Save(transaction).Error
}

func (tr *TransactionRepo) CreateAccount(account *models.Account) error {
	return tr.db.Save(account).Error
}

func (tr *TransactionRepo) FetchAccount(accountId int64) (*models.Account, error) {
	var account models.Account
	err := tr.db.Model(&models.Account{}).Where("id = ?", accountId).First(&account).Error

	return &account, err
}
