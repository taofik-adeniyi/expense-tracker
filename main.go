package main

import (
	"fmt"
	"log"
	"os"
)

const dbFileName = "expense-db.json"
const balanceFileName = "balance-db.json"

func main() {
	flags := os.Args
	if len(flags) < 2 {
		log.Fatalf("no command passed: check the readme for valid commands")
	}

	createFileIfNotExists(dbFileName)
	createFileIfNotExists(balanceFileName)
	SetDefaultBalance()

	switch flags[1] {
	case "add":
		HandleAddition(flags[0:])
	case "update":
		HandleUpdate(flags[0:])
	case "list":
		HandleList(flags[0:])
	case "summary":
		HandleSummary(flags[0:])
	case "delete":
		HandleDelete(flags[0:])
	case "add-category":
		HandleCategory(flags[0:])
	case "set-budget":
		HandleBudget(flags[0:])
	case "export":
		HandleExport(flags[0:])
	case "--help", "help":
		fmt.Println("Help command read the README file")
	case "--version", "-V":
		fmt.Println("expense-tracker@0.0.1")
	default:
		fmt.Println("Invalid command passed supported commands are <add, list, update, delete, summary, set-budget, export, help, add-category>")
	}

}
