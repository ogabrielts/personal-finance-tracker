package main

import (
	"database/sql"
	"expense-api/config"
	"expense-api/internal/user"
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

	userModule := user.NewModule(db)

	mux := http.NewServeMux()

	userModule.UserRouter(mux)

	log.Printf("Server is running...")
	log.Fatal(http.ListenAndServe(config.Address, mux))

}