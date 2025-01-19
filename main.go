package main

import (
	"fmt"
	"os"
	// "github.com/taofik-adeniyi/expense-tracker/utils"
)

const dbFileName = "expense-db.json"
const balanceFileName = "balance-db.json"

func main() {
	createFileIfNotExists(dbFileName)
	createFileIfNotExists(balanceFileName)
	SetDefaultBalance()
	flags := os.Args

	// fflags := utils.GetUserInput()
	if len(flags) < 2 {
		fmt.Println("no command passed")
		os.Exit(1)
	}

	switch flags[1] {
	case "add":
		fmt.Println("adding...")
		HandleAddition(flags[0:])
	case "update":
		//
	case "list":
	case "delete":
	default:
		fmt.Println("Invalid command passed supported commands are <add, list, update, delete, summary>")
	}

}

// Add error handling to handle invalid inputs and edge cases (e.g. negative amounts, non-existent expense IDs, etc).
