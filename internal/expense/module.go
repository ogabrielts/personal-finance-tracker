package expense

import (
	"database/sql"
	"net/http"
)

type expenseModule struct {
	handler *expenseHandler
}

func NewModule(db *sql.DB) *expenseModule {
	repo := newDB(db)
	service := newService(repo)
	handler := newHandler(service)

	return &expenseModule{
		handler: handler,
	}
}

func (em *expenseModule) ExpenseRouter(mux *http.ServeMux, auth func(http.HandlerFunc) http.HandlerFunc) {
	mux.HandleFunc("POST /v1/expense", auth(em.handler.CreateExpenseHandler))
	mux.HandleFunc("GET /v1/expense", auth(em.handler.GetAllExpensesHandler))
	mux.HandleFunc("GET /v1/expense/{id}", auth(em.handler.GetExpenseByIdHandler))
	mux.HandleFunc("PUT /v1/expense/{id}", auth(em.handler.UpdateExpenseHandler))
	mux.HandleFunc("DELETE /v1/expense/{id}", auth(em.handler.DeleteExpenseHandler))
}
