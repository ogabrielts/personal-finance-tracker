package main

import (
	"database/sql"
	"expense-api/api/config"
	"expense-api/api/internal/auth"
	"expense-api/api/internal/expense"
	"expense-api/api/internal/income"
	"expense-api/api/internal/middleware"
	"expense-api/api/internal/reports"
	"expense-api/api/internal/user"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(config.Driver, config.DB)
	if err != nil {
		log.Fatal(err)
	}

	authModule := auth.NewModule(db)
	userModule := user.NewModule(db)
	expenseModule := expense.NewModule(db)
	incomeModule := income.NewModule(db)
	reportsModule := reports.NewModule(db)

	mux := http.NewServeMux()

	authModule.AuthRouter(mux)
	userModule.UserRouter(mux)
	expenseModule.ExpenseRouter(mux, middleware.Authenticate)
	incomeModule.IncomeRouter(mux, middleware.Authenticate)
	reportsModule.ReportsRouter(mux, middleware.Authenticate)

	log.Printf("Server is running...")
	log.Fatal(http.ListenAndServe(config.Address, mux))

}
