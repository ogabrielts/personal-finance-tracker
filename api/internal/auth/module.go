package auth

import (
	"database/sql"
	"net/http"
)

type authModule struct {
	handler *authHandler
}

func NewModule(db *sql.DB) *authModule {
	repo := newDB(db)
	service := newService(repo)
	handler := newHandler(service)

	return &authModule{
		handler: handler,
	}
}

func (am *authModule) AuthRouter(mux *http.ServeMux) {
	mux.HandleFunc("POST /v1/login", am.handler.LogInHandler)
	mux.HandleFunc("POST /v1/logout", am.handler.LogOutHandler)
}
