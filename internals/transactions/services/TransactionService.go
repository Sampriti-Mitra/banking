package services

import (
	"github.com/Sampriti-Mitra/transactions/internals/transactions/models"
	"github.com/Sampriti-Mitra/transactions/internals/transactions/repo"
)

type TransactionService struct {
	tRepo repo.TransactionRepo
}

func NewTransactionService(tRepo repo.TransactionRepo) *TransactionService {
	return &TransactionService{tRepo: tRepo}
}

func (ts *TransactionService) CreateAccount(acc *models.Account) (*models.Account, error) {
	err := ts.tRepo.CreateAccount(acc)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (ts *TransactionService) FetchAccount(accId int64) (*models.Account, error) {
	acc, err := ts.tRepo.FetchAccount(accId)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func (ts *TransactionService) CreateTransaction(txn *models.Transaction) (*models.Transaction, error) {
	err := txn.ValidateTransaction()
	if err != nil {
		return nil, err
	}
	err = ts.tRepo.CreateTransaction(txn)
	if err != nil {
		return nil, err
	}
	return txn, nil
}
