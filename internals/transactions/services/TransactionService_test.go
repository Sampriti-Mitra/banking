package services

import (
	"reflect"
	"testing"

	"github.com/Sampriti-Mitra/transactions/internals/transactions/models"
	"github.com/Sampriti-Mitra/transactions/internals/transactions/repo"
)

func TestNewTransactionService(t *testing.T) {
	type args struct {
		tRepo repo.TransactionRepo
	}
	tests := []struct {
		name string
		args args
		want *TransactionService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransactionService(tt.args.tRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransactionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionService_CreateAccount(t *testing.T) {
	type fields struct {
		tRepo repo.TransactionRepo
	}
	type args struct {
		acc *models.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TransactionService{
				tRepo: tt.fields.tRepo,
			}
			got, err := ts.CreateAccount(tt.args.acc)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionService.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionService.CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionService_FetchAccount(t *testing.T) {
	type fields struct {
		tRepo repo.TransactionRepo
	}
	type args struct {
		accId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TransactionService{
				tRepo: tt.fields.tRepo,
			}
			got, err := ts.FetchAccount(tt.args.accId)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionService.FetchAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionService.FetchAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionService_CreateTransaction(t *testing.T) {
	type fields struct {
		tRepo repo.TransactionRepo
	}
	type args struct {
		txn *models.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TransactionService{
				tRepo: tt.fields.tRepo,
			}
			got, err := ts.CreateTransaction(tt.args.txn)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionService.CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionService.CreateTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
