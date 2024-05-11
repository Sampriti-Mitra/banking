package main

import (
	"log"
	"net/http"

	"github.com/Sampriti-Mitra/transactions/internals/transactions/controllers"
	"github.com/Sampriti-Mitra/transactions/internals/transactions/repo"
	"github.com/Sampriti-Mitra/transactions/internals/transactions/services"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(localhost:3306)/bank?parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	txnRepo := repo.NewTransactionRepo(db)
	txnSvc := services.NewTransactionService(*txnRepo)
	txnCntrl := controllers.NewTransactionController(txnSvc)

	r := mux.NewRouter()
	r.HandleFunc("/accounts", txnCntrl.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{accountId}", txnCntrl.GetAccount).Methods("GET")
	r.HandleFunc("/transactions", txnCntrl.CreateTransaction).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
