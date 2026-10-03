package main

import (
	"expense-api/tui/internal/components"
	"expense-api/tui/internal/utils"
	"fmt"
	"log"
)

func main() {
	email, err := utils.Input("Email: ")
	if err != nil {
		log.Fatal(err)
	}

	password, err := utils.Input("Password: ")
	if err != nil {
		log.Fatal(err)
	}

	client, err := components.SignIn(email, password)
	if err != nil {
		log.Fatal(err)
	}

	for {
		cmd, err := utils.Input("pet= ")
		if err != nil {
			log.Fatal(err)
		}

		args, err := utils.Tokenize(cmd)
		if err != nil {
			log.Fatal(err)
		}

		switch args[0] {
		case "exit":
			return
	
		case "create":
			if err := components.Create(client, args[1:]); err != nil {
				log.Fatal(err)
			}

		case "edit":
			if err := components.Edit(client, args[1:]); err != nil {
				log.Fatal(err)
			}

		case "delete":
			if err := components.Delete(client, args[1:]); err != nil {
				log.Fatal(err)
			}

		case "list":
			if err := components.List(client, args[1:]); err != nil {
				log.Fatal(err)
			}

		case "report":
			if err := components.Report(client, args[1:]); err != nil {
				log.Fatal(err)
			}

		default:
			fmt.Println("Unknown command")
		}
	}
}
