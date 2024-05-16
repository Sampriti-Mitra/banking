package services

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Sampriti-Mitra/transactions/internals/transactions/mocks"
	"github.com/Sampriti-Mitra/transactions/internals/transactions/models"
	"github.com/Sampriti-Mitra/transactions/internals/transactions/repo"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestNewTransactionService(t *testing.T) {
	type args struct {
		tRepo repo.ITransactionRepo
	}
	tests := []struct {
		name string
		args args
		want *TransactionService
	}{
		{
			name: "Create Transaction Service",
			args: args{tRepo: nil},
			want: &TransactionService{tRepo: nil},
		},
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

	mockRepo := mocks.NewMockITransactionRepo(gomock.NewController(t))
	type fields struct {
		tRepo repo.ITransactionRepo
	}
	type args struct {
		acc *models.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		f       func(acc *models.Account)
		want    *models.Account
		wantErr bool
	}{
		{
			name: "Create Account Success",
			fields: fields{
				tRepo: mockRepo,
			},
			args: args{
				acc: &models.Account{DocumentId: "qbdwoq123456"},
			},
			f: func(acc *models.Account) {
				mockRepo.EXPECT().CreateAccount(acc).Return(nil)
			},
			want: &models.Account{DocumentId: "qbdwoq123456"},
		},
		{
			name: "Create Account Missing Document Id",
			fields: fields{
				tRepo: mockRepo,
			},
			f: func(acc *models.Account) {
				mockRepo.EXPECT().CreateAccount(acc).Return(errors.New("document id missing"))
			},
			args: args{
				acc: &models.Account{DocumentId: "qbdwoq123456"},
			},
			wantErr: true,
		},
		{
			name: "Create Account DB duplicate key error",
			fields: fields{
				tRepo: mockRepo,
			},
			f: func(acc *models.Account) {
				mockRepo.EXPECT().CreateAccount(acc).Return(errors.New("Duplicate key"))
			},
			args: args{
				acc: &models.Account{DocumentId: "qbdwoq123456"},
			},
			wantErr: true,
		},
		{
			name: "Create Account DB error",
			fields: fields{
				tRepo: mockRepo,
			},
			f: func(acc *models.Account) {
				mockRepo.EXPECT().CreateAccount(acc).Return(errors.New("random error"))
			},
			args: args{
				acc: &models.Account{DocumentId: "qbdwoq123456"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TransactionService{
				tRepo: tt.fields.tRepo,
			}

			tt.f(tt.args.acc)

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
	mockRepo := mocks.NewMockITransactionRepo(gomock.NewController(t))
	type fields struct {
		tRepo repo.ITransactionRepo
	}
	type args struct {
		accId int64
	}
	tests := []struct {
		name    string
		fields  fields
		f       func()
		args    args
		want    *models.Account
		wantErr bool
	}{
		{
			name:   "Missing account id",
			fields: fields{tRepo: mockRepo},
			args:   args{accId: 1},
			f: func() {
				mockRepo.EXPECT().FetchAccount(int64(1)).Return(nil, gorm.ErrRecordNotFound)
			},
			wantErr: true,
		},
		{
			name:   "Fetch account db error",
			fields: fields{tRepo: mockRepo},
			args:   args{accId: 1},
			f: func() {
				mockRepo.EXPECT().FetchAccount(int64(1)).Return(nil, errors.New("account not found"))
			},
			wantErr: true,
		},
		{
			name:   "Valid account id",
			fields: fields{tRepo: mockRepo},
			args:   args{accId: 1},
			f: func() {
				mockRepo.EXPECT().FetchAccount(int64(1)).Return(&models.Account{ID: 1}, nil)
			},
			want:    &models.Account{ID: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TransactionService{
				tRepo: tt.fields.tRepo,
			}
			tt.f()
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
	ctrl := gomock.NewController(t)
	mockRepo := mocks.NewMockITransactionRepo(ctrl)

	testTxnModel := &models.Transaction{
		AccountID:     1,
		OperationType: 1,
		Amount:        -120,
		ID:            1,
		EventDate:     time.Now(),
		Balance:       -120,
	}
	type fields struct {
		tRepo repo.ITransactionRepo
	}
	type args struct {
		txn *models.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		f       func()
		want    *models.Transaction
		wantErr bool
	}{
		{
			name: "create valid transaction",
			fields: fields{
				tRepo: mockRepo,
			},
			args: args{
				txn: testTxnModel,
			},
			f: func() {
				mockRepo.EXPECT().FetchAccount(int64(1)).Return(nil, nil)
				mockRepo.EXPECT().CreateTransaction(testTxnModel).Return(nil)
			},
			want:    testTxnModel,
			wantErr: false,
		},
		{
			name: "create transaction with invalid accountId",
			fields: fields{
				tRepo: mockRepo,
			},
			args: args{
				txn: testTxnModel,
			},
			f: func() {
				mockRepo.EXPECT().FetchAccount(int64(1)).Return(nil, errors.New("account id not found"))
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "create transaction with invalid operation type",
			fields: fields{
				tRepo: mockRepo,
			},
			args: args{
				txn: &models.Transaction{
					AccountID:     1,
					OperationType: 5, // invalid op type
					Amount:        -120,
					ID:            1,
					EventDate:     time.Now(),
				},
			},
			f: func() {

			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "create transaction with invalid amount for operation",
			fields: fields{
				tRepo: mockRepo,
			},
			args: args{
				txn: &models.Transaction{
					AccountID:     1,
					OperationType: 4, // invalid op type
					Amount:        -120.0,
					ID:            1,
					EventDate:     time.Now(),
				},
			},
			f: func() {
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TransactionService{
				tRepo: tt.fields.tRepo,
			}

			tt.f()

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
