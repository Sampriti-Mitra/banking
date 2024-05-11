package controllers

import (
	"net/http"

	"github.com/Sampriti-Mitra/transactions/internals/transactions/services"
)

type TransactionController struct {
	ts *services.TransactionService
}

func NewTransactionController(ts *services.TransactionService) *TransactionController {
	return &TransactionController{ts: ts}
}

func (tc *TransactionController) CreateAccount(w http.ResponseWriter, r *http.Request) {

}

func (tc *TransactionController) GetAccount(w http.ResponseWriter, r *http.Request) {

}

func (tc *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {

}
