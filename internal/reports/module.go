package reports

import (
	"database/sql"
	"net/http"
)

type reportsModule struct {
	handler *reportsHandler
}

func NewModule(db *sql.DB) *reportsModule {
	repo := newDB(db)
	service := newService(repo)
	handler := newHandler(service)

	return &reportsModule{
		handler: handler,
	}
}

func (rm *reportsModule) ReportsRouter(mux *http.ServeMux, auth func(http.HandlerFunc) http.HandlerFunc) {
	mux.HandleFunc("POST /v1/reports/full", auth(rm.handler.allReportHandler))
	mux.HandleFunc("POST /v1/reports/monthly", auth(rm.handler.monthlyReportHandler))
}
