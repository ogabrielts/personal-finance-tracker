package main

import (
	"expense-api/tui/internal/components"
	"expense-api/tui/internal/utils"
	"fmt"
	"log"

	"golang.org/x/term"
)

func main() {
	fmt.Println("Enter with your account or create a new one")
	fmt.Println("'enter' to log in | 'create' to create a new account")
	start, err := utils.Input(">>> ")
	if err != nil {
		log.Fatal(err)
	}

	switch start {
	case "create":
		name, err := utils.Input("Name: ")
		if err != nil {
			log.Fatal(err)
		}
		email, err := utils.Input("Email: ")
		if err != nil {
			log.Fatal(err)
		}
	
		fmt.Print("Password: ")
		password, err := term.ReadPassword(0)
		if err != nil {
			log.Fatal(err)
		}

		address, err := utils.Input("Address: ")
		if err != nil {
			log.Fatal(err)
		}

		if err := components.SignUp(name, email, string(password), address); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nUser successfully created.\n\n")
		fmt.Println("------------------------------")

	case "enter":
		break
	}

	email, err := utils.Input("Email: ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Password: ")
	password, err := term.ReadPassword(0)
	if err != nil {
		log.Fatal(err)
	}

	client, err := components.SignIn(email, string(password))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nUser successfully logged in.\n\n")
	fmt.Println("------------------------------")

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

		case "filter":
			if err := components.Filter(client, args[1:]); err != nil {
				log.Fatal(err)
			}

		default:
			fmt.Println("Unknown command")
		}
	}
}
