package services

import (
	"errors"
	"strings"
	"time"

	"github.com/Sampriti-Mitra/transactions/internals/transactions/models"
	"github.com/Sampriti-Mitra/transactions/internals/transactions/repo"
	"github.com/Sampriti-Mitra/transactions/internals/utils"
	"gorm.io/gorm"
)

type ITransactionService interface {
	CreateAccount(acc *models.Account) (*models.Account, *utils.CustError)
	FetchAccount(accId int64) (*models.Account, *utils.CustError)
	CreateTransaction(txn *models.Transaction) (*models.Transaction, *utils.CustError)
}

type TransactionService struct {
	tRepo repo.ITransactionRepo
}

func NewTransactionService(tRepo repo.ITransactionRepo) *TransactionService {
	return &TransactionService{tRepo: tRepo}
}

func (ts *TransactionService) CreateAccount(acc *models.Account) (*models.Account, *utils.CustError) {
	err := acc.Validate()
	if err != nil {
		return nil, utils.NewCustError(errors.New("bad request error"), 400)
	}
	err = ts.tRepo.CreateAccount(acc)
	if err != nil {
		if err == gorm.ErrDuplicatedKey || strings.Contains(err.Error(), "Duplicate") {
			return nil, utils.NewCustError(errors.New("bad request error"), 400)
		}
		return nil, utils.NewCustError(errors.New("internal server error"), 500)
	}
	return acc, nil
}

func (ts *TransactionService) FetchAccount(accId int64) (*models.Account, *utils.CustError) {
	acc, err := ts.tRepo.FetchAccount(accId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewCustError(errors.New("account not found"), 400)
		}
		return nil, utils.NewCustError(errors.New("internal server error"), 500)
	}
	return acc, nil
}

func (ts *TransactionService) CreateTransaction(txn *models.Transaction) (*models.Transaction, *utils.CustError) {
	err := txn.ValidateTransaction()
	if err != nil {
		return nil, utils.NewCustError(errors.New("bad request error"), 400)
	}

	_, custErr := ts.FetchAccount(txn.AccountID)
	if custErr != nil {
		return nil, custErr
	}

	txn.EventDate = time.Now()

	err = ts.tRepo.CreateTransaction(txn)
	if err != nil {
		return nil, utils.NewCustError(errors.New("internal server error"), 500)
	}
	return txn, nil
}
