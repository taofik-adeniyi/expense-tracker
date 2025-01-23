package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var categories []string

func HandleCategory(flags []string) {
	if len(flags) != 3 {
		log.Fatal("Invalid command")
	}
	if flags[1] != "add-category" {
		log.Fatal("Invalid command")
	}
	category_name := flags[2]
	categories = append(categories, category_name)
	fmt.Printf("Category: %v added successfully\n", category_name)
}

func HandleExport(flags []string) {
	if len(flags) != 3 {
		log.Fatal("Invalid export command\nExpense-tracker export <filename.csv>")
	}
	var fileName = flags[2]
	data := strings.Split(fileName, ".")
	ext := data[1]
	if ext != "csv" {
		log.Fatal("Invalid file name format provide a .csv file format")
	}
	fmt.Print(flags)

	ExportExpensesToCsv(fileName)
}

func HandleBudget(flags []string) {
	if len(flags) != 6 {
		log.Fatal("Invalid command")
	}
	if flags[1] != "set-budget" {
		log.Fatal("invalid command")
	}
	month, err := strconv.Atoi(flags[3])
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	amount, err := strconv.Atoi(flags[5])
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	res, err := SetBudget(month, amount)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Printf("Budget of month %v set to: $%d \n", monthsOfYear[month], res)
}
func HandleUpdate(f []string) {
	if len(f) < 5 || len(f) > 9 {
		log.Fatal("update command error")
	}

	id, err := strconv.Atoi(f[2])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var descriptionValue = f[4]
	amountValue, err := strconv.Atoi(f[6])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	var categoryValue = f[8]

	res, err := Update(id, descriptionValue, amountValue, categoryValue)
	if err != nil {
		log.Fatalf("Error Updating: %v", err)
	}
	fmt.Println(res)
}
func HandleDelete(f []string) {
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

func SummaryCategory(flags []string) {
	expenses, err := ListExpenses()
	if err != nil {
		log.Fatal(err)
	}
	res := expenses.SummaryByCategory(flags[1])
	fmt.Println(res)
}
func SummaryMonth(flags []string) {
	monthId, err := strconv.Atoi(flags[1])
	if err != nil {
		log.Fatalf("Error: month value should be in the range 1-12 not more not less\nWhere 1-12 represents january to december\n%v", err)
	}

	expenses, err := ListExpenses()
	if err != nil {
		log.Fatal(err)
	}

	if monthId >= 1 && monthId <= 12 {
		result := expenses.Summary(monthId)
		fmt.Println(result)
		return
	}
	fmt.Println("Invalid month value passed")
}

func HandleSummary(f []string) {
	if len(f) < 2 || len(f) > 4 {
		log.Fatalf("Invalid summary command")
	}

	if len(f) == 2 {
		expenses, err := ListExpenses()
		if err != nil {
			log.Fatal(err)
		}

		result := expenses.Summary()
		fmt.Println(result)
		return
	}

	switch f[2] {
	case "--month":
		SummaryMonth(f[2:])
		return
	case "--category":
		SummaryCategory(f[2:])
		return
	}

}

func HandleAddition(f []string) {
	if len(f) < 6 && len(f) > 8 {
		log.Fatal("Invalid command")
	}
	var description string
	var amount int
	var month int

	var secondCommand = f[2]
	var thirdCommand = f[4]
	var sixthCommand string
	if len(f) >= 8 {
		sixthCommand = f[6]
	}

	if secondCommand == "--description" {
		description = f[3]
	}
	fmt.Println("month", month)
	if thirdCommand == "--amount" {
		n, err := strconv.Atoi(f[5])
		if err != nil {
			log.Fatalf("Error: parsing %v\nProvide a valid number greater than or equals 0.\n", f[5])
		}
		amount = n
	}

	if sixthCommand == "--month" {
		n, err := strconv.Atoi(f[7])
		if err != nil {
			log.Fatalf("Error: parsing %v\nProvide a valid number greater than or equals 0.\n", f[6])
		}
		month = n
	}

	if month > 0 {
		added, err := Add(description, amount, month)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Printf("# Expense added successfully (ID: %v)\n", added)
	} else {
		added, err := Add(description, amount)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Printf("# Expense added successfully (ID: %v)\n", added)
	}

}
