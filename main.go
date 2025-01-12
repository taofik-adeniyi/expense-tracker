package main

import (
	"fmt"
	"os"
	// "github.com/taofik-adeniyi/expense-tracker/utils"
)

func main() {
	flags := os.Args
	fmt.Println("len", len(flags))

	// fflags := utils.GetUserInput()
	if len(flags) < 2 {
		fmt.Println("no command passed")
		os.Exit(1)
	}

	switch flags[1] {
	case "add":
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
