package repo

import (
	"math"

	"github.com/Sampriti-Mitra/transactions/internals/transactions/models"
	"gorm.io/gorm"
)

type ITransactionRepo interface {
	CreateTransaction(transaction *models.Transaction) error
	CreateAccount(account *models.Account) error
	FetchAccount(accountId int64) (*models.Account, error)
	CreateTransactionWithUpdatedBalance(transaction *models.Transaction) error
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

func (tr *TransactionRepo) CreateTransactionWithUpdatedBalance(transaction *models.Transaction) error {
	err := tr.db.Transaction(func(tx *gorm.DB) error {

		// fetch the rows with negative balances
		// decrement the current balance to satisfy the negative balances
		// update the table rows
		// lock rows select for update

		var validTxns []models.Transaction

		// select ... for update
		err := tx.Model(&models.Transaction{}).Where("balance < 0 and account_id = ?", transaction.AccountID).Find(&validTxns).Error
		if err != nil {
			return err
		}

		currBal := transaction.Amount
		txnIdsWithBalZero := []uint{}

		for _, validTxn := range validTxns {
			if currBal == 0 {
				break
			}
			if math.Abs(validTxn.Balance) < currBal {
				currBal = currBal + validTxn.Balance
				txnIdsWithBalZero = append(txnIdsWithBalZero, validTxn.ID)
			} else {
				validTxn.Balance = validTxn.Balance + currBal
				currBal = 0
				// update
				err = tx.Model(&models.Transaction{}).Where("id = ?", validTxn.ID).UpdateColumn("balance", validTxn.Balance).Error
				if err != nil {
					return err
				}
			}
		}

		transaction.Balance = currBal

		err = tx.Model(&models.Transaction{}).Where("id in ?", txnIdsWithBalZero).UpdateColumn("balance", 0).Error
		if err != nil {
			return err
		}

		return tr.db.Save(transaction).Error

	})
	return err
}

func (tr *TransactionRepo) CreateAccount(account *models.Account) error {
	return tr.db.Save(account).Error
}

func (tr *TransactionRepo) FetchAccount(accountId int64) (*models.Account, error) {
	var account models.Account
	err := tr.db.Model(&models.Account{}).Where("id = ?", accountId).First(&account).Error

	return &account, err
}
