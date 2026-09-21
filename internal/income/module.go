package income

import (
	"database/sql"
	"net/http"
)

type incomeModule struct {
	handler *incomeHandler 
}

func NewModule(db *sql.DB) *incomeModule {
	repo := newDB(db)
	service := newService(repo)
	handler := newHandler(service)

	return &incomeModule{
		handler: handler,
	}
}

func (im *incomeModule) IncomeRouter(mux *http.ServeMux, auth func(http.HandlerFunc) http.HandlerFunc) {
	mux.HandleFunc("POST /v1/income", auth(im.handler.CreateIncomeHandler))
	mux.HandleFunc("GET /v1/income", auth(im.handler.GetAllIncomesHandler))
	mux.HandleFunc("GET /v1/income/{id}", auth(im.handler.GetIncomeByIdHandler))
	mux.HandleFunc("PUT /v1/income/{id}", auth(im.handler.UpdateIncomeHandler))
	mux.HandleFunc("DELETE /v1/income/{id}", auth(im.handler.DeleteIncomeHandler))
}
