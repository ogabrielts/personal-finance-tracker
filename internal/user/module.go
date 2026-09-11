package user

import (
	"database/sql"
	"net/http"
)

type userModule struct {
	handler *userHandler
}

func NewModule(db *sql.DB) *userModule {
	repo := newDB(db)
	service := newService(repo)
	handler := newHandler(service)

	return &userModule{
		handler: handler,
	}
}

func (um *userModule) UserRouter(mux *http.ServeMux) {
	mux.HandleFunc("POST /v1/user", um.handler.CreateUserHandler)
	mux.HandleFunc("GET /v1/user/{id}", um.handler.GetUserByIdHandler)
	mux.HandleFunc("PUT /v1/user/{id}", um.handler.UpdateUserHandler)
	mux.HandleFunc("DELETE /v1/user/{id}", um.handler.DeleteUserHandler)
}