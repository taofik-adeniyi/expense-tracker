package main

import (
	"fmt"
	"os"
)

const dbFileName = "expense-db.json"
const balanceFileName = "balance-db.json"

func main() {
	createFileIfNotExists(dbFileName)
	createFileIfNotExists(balanceFileName)
	SetDefaultBalance()
	flags := os.Args

	if len(flags) < 2 {
		fmt.Println("no command passed")
		os.Exit(1)
	}

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
		break
	case "set-budget":
		HandleBudget(flags[0:])
	case "export":
		HandleExport(flags[0:])
	case "--help":
		fmt.Println("Help command read the README file")
		break
	default:
		fmt.Println("Invalid command passed supported commands are <add, list, update, delete, summary>")
	}

}
