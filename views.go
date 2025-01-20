package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func HandleUpdate(f []string) {
	// fmt.Println(len(f))
	if len(f) < 5 || len(f) > 9 {
		log.Fatal("update command error")
	}
	// var optionalArgs = f[5:9]
	// fmt.Println(f[5]) // --amount
	// fmt.Println("optionalArgs", optionalArgs)
	// expense-tracker update id --description "" --amount "" --category "" //9
	// result, err := Update()
	id, err := strconv.Atoi(f[2])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// var descriptionFlag = f[3]
	var descriptionValue = f[4]
	// var amountFlag = f[5]
	amountValue, err := strconv.Atoi(f[6])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// var categoryFlag = f[7]
	var categoryValue = f[8]
	// fmt.Println("id", id)
	// fmt.Println("descriptionFlag:descriptionValue", descriptionFlag, descriptionValue)
	// fmt.Println("amountFlag", amountFlag, amountValue)
	// fmt.Println("categoryFlag", categoryFlag, categoryValue)

	res, err := Update(id, descriptionValue, amountValue, categoryValue)
	if err != nil {
		log.Fatalf("Error Updating: %v", err)
	}
	fmt.Println(res)
}
func HandleDelete(f []string) {
	// check the length of the commands after the executable name
	// expense-tracker delete --id 2
	// Expense deleted successfully
	if len(f) != 4 {
		log.Fatal("Invalid delete command: the valid command is expense-tracker delete --id <id>")
	}
	if f[2] == "--id" {
		id, err := strconv.Atoi(f[3])
		if err != nil {
			log.Fatalf("invalid format type of id: %v", err)
		}
		res, err := Delete(id)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Printf("# Expense with id of: %d deleted successfully\n", res)
	}
}

func HandleList(f []string) {
	fmt.Println("expense-tracker list")
	if len(f) != 2 {
		fmt.Println("Invalid list comamnd")
		os.Exit(1)
	}
	expenses, err := ListExpenses()
	if err != nil {
		log.Fatal(err)
	}
	expenses.formatExpenses()
}

func HandleSummary(f []string) {
	if len(f) < 2 || len(f) > 4 {
		log.Fatalf("Invalid summary command")
	}
	var monthFlag = f[2]

	monthValue, err := strconv.Atoi(f[3])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	expenses, err := ListExpenses()
	if err != nil {
		log.Fatal(err)
	}

	if monthValue > 0 && monthFlag == "--month" {
		result := expenses.Summary(monthValue)
		fmt.Println(result)
		return
	}

	result := expenses.Summary()
	fmt.Println(result)
}

func HandleAddition(f []string) {
	fmt.Println("expense-tracker add")
	if len(f) != 8 {
		os.Exit(1)
	}
	var description string
	var amount int
	var month int

	var second = f[2]
	var fourth = f[4]
	var sixth = f[6]

	if second == "--description" {
		description = f[3]
	}
	if fourth == "--amount" {
		n, err := strconv.Atoi(f[5])
		if err != nil {
			fmt.Printf("Error: parsing %v\nProvide a valid number greater than or equals 0.\n", f[5])
			return
		}
		amount = n
	}
	if sixth == "--month" {
		n, err := strconv.Atoi(f[7])
		if err != nil {
			fmt.Printf("Error: parsing %v\nProvide a valid number greater than or equals 0.\n", f[6])
			return
		}
		month = n
	}

	added, err := Add(description, amount, month)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("# Expense added successfully (ID: %v)\n", added)
}
