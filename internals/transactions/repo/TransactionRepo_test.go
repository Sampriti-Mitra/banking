package repo

import (
	"fmt"
	"testing"
	"time"

	"github.com/Sampriti-Mitra/transactions/internals/transactions/models"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestTransactionRepo_CreateTransaction(t *testing.T) {

	db := dbInit()

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		transaction *models.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Create transaction",
			fields: fields{db: db},
			args: args{
				transaction: &models.Transaction{
					AccountID:     1,
					Amount:        -100,
					OperationType: 1,
					Balance:       -100,
					EventDate:     time.Now(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TransactionRepo{
				db: tt.fields.db,
			}
			if err := tr.CreateTransaction(tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("TransactionRepo.CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransactionRepo_CreateTransactionWithUpdatedBalance(t *testing.T) {
	db := dbInit()
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		transaction *models.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		f       func()
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Create transaction",
			fields: fields{db: db},
			args: args{
				transaction: &models.Transaction{
					AccountID:     1,
					Amount:        100,
					OperationType: 4,
					EventDate:     time.Now(),
				},
			},
			f: func() {
				transaction := &models.Transaction{
					AccountID:     1,
					Amount:        -100,
					OperationType: 1,
					EventDate:     time.Now(),
					Balance:       -100,
				}
				err := db.Save(transaction).Error
				fmt.Print(err)

			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TransactionRepo{
				db: tt.fields.db,
			}

			tt.f()

			if err := tr.CreateTransactionWithUpdatedBalance(tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("TransactionRepo.CreateTransactionWithUpdatedBalance() error = %v, wantErr %v", err, tt.wantErr)
			}

			var tx models.Transaction
			db.Model(&models.Transaction{}).Where("account_id=1 and id=2").Find(&tx)

			assert.Equal(t, tx.Balance, 0.0)

		})
	}
}

func dbInit() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		"root",
		"",
		"localhost",
		"3306",
		"bank")

	var db *gorm.DB
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
		return nil
	}

	db.Exec("truncate table transactions")

	return db
}
