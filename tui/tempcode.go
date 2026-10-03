package main

import (
	"expense-api/tui/internal"
	"flag"
	"fmt"
	"log"
	"os"
)

func tempcode() {
	args := os.Args
	
	signUpCmd := flag.NewFlagSet("sign-up", flag.ExitOnError)
	signUpName := signUpCmd.String("name", "", "Name")
	signUpEmail := signUpCmd.String("email", "", "Email")
	signUpPassword := signUpCmd.String("password", "", "Password")
	signUpAddress := signUpCmd.String("address", "", "Address")

	signInCmd := flag.NewFlagSet("sign-in", flag.ExitOnError)
	signInEmail := signInCmd.String("email", "", "Email")
	signInPassword := signInCmd.String("password", "", "Password")

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createName := createCmd.String("name", "", "A name description")
	createValue := createCmd.Float64("value", 0.0, "The value")
	createCategory := createCmd.String("category", "", "A category")
	createAccount := createCmd.String("account-type", "", "Account type")

	editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
	editId := editCmd.String("id", "", "The reference code")
	editName := editCmd.String("name", "", "A name description")
	editValue := editCmd.Float64("value", 0.0, "The value")
	editCategory := editCmd.String("category", "", "A category")
	editAccount := editCmd.String("account-type", "", "Account type")
	
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCmd.String("id", "", "The reference code")
	deleteAccount := deleteCmd.String("account-type", "", "Account type")
	
	

	switch (args[1]) {
	case "sign-in":
		signInCmd.Parse(args[2:])

		if err := internal.SignIn(*signInEmail, *signInPassword); err != nil {
			log.Fatalln(err)
		}

	case "sign-up":
		signUpCmd.Parse(args[2:])

		if err := internal.SignUp(*signUpName, *signUpEmail, *signUpPassword, *signUpAddress); err != nil {
			log.Fatalln(err)
		}

	case "create":
		createCmd.Parse(args[2:])

		if err := internal.Create(*createAccount, *createName, *createCategory, *createValue); err != nil {
			log.Fatalln(err)
		}

	case "edit":
		editCmd.Parse(args[2:])

		if err := internal.Edit(*editAccount, *editId, *editName, *editCategory, *editValue); err != nil {
			log.Fatalln(err)
		}

	case "delete":
		deleteCmd.Parse(args[2:])

		if err := internal.Delete(*deleteAccount, *deleteId); err != nil {
			log.Fatalln(err)
		}

	case "list":
		fmt.Println("User is trying to list all or just one income or expense")

	case  "report":
		fmt.Println("User is trying to see the monthly report or full report")

	default:
		fmt.Println("Unknown command")
	}
}
