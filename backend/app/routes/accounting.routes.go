package routes

import (
	"trenova/app/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RevenueCodeRoutes(r *mux.Router, db *gorm.DB) {
	revCodeRouter := r.PathPrefix("/revenue-codes").Subrouter()

	revCodeRouter.HandleFunc("/", handlers.GetRevenueCodes(db)).Methods("GET")
	revCodeRouter.HandleFunc("/", handlers.CreateRevenueCode(db)).Methods("POST")
	revCodeRouter.HandleFunc("/{revenueCodeID}/", handlers.GetRevenueCodeByID(db)).Methods("GET")
	revCodeRouter.HandleFunc("/{revenueCodeID}/", handlers.UpdateRevenueCode(db)).Methods("PUT")
}

func AccountingControlRoutes(r *mux.Router, db *gorm.DB) {
	acRouter := r.PathPrefix("/accounting-control").Subrouter()

	acRouter.HandleFunc("/", handlers.GetAccountingControl(db)).Methods("GET")
	acRouter.HandleFunc("/", handlers.UpdateAccountingControl(db)).Methods("PUT")
}

func BillingControlRoutes(r *mux.Router, db *gorm.DB) {
	bcRouter := r.PathPrefix("/billing-control").Subrouter()

	bcRouter.HandleFunc("/", handlers.GetBillingControl(db)).Methods("GET")
	bcRouter.HandleFunc("/", handlers.UpdateBillingControl(db)).Methods("PUT")
}

func InvoiceControlRoutes(r *mux.Router, db *gorm.DB) {
	icRouter := r.PathPrefix("/invoice-control").Subrouter()

	icRouter.HandleFunc("/", handlers.GetInvoiceControl(db)).Methods("GET")
	icRouter.HandleFunc("/", handlers.UpdateInvoiceControl(db)).Methods("PUT")
}