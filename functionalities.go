package main

import (
	"fmt"
	"time"
)

type Expense struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Category    string    `json:"category"`
}

var budget map[int]int = make(map[int]int, 0) // map month to the months budget

func Add(description string, amount int) int {
	fmt.Println("expense-tracker", description, amount)
	// check budget balance
	// if amount is greater than budget balance return error an terminate function execution
	// # Expense added successfully (ID: 2)
	return 0
}

func List() []Expense {
	// 	# ID  Date       Description  Amount
	// # 1   2024-08-06  Lunch        $20
	// # 2   2024-08-06  Dinner       $10
	return []Expense{}
}

func Summary(month *int) {
	// # Total expenses: $30
}

func Delete() int {
	// # Expense deleted successfully
	return 0
}
func Update(id int) {

}
func SetBudget(month int, amount int) (int, error) {
	// 1->january,2,3,4,5,6,7,8,9,10,11,12
	// the month cant be more or less than the numbers 1-12
	// set budget for each month
	if month >= 1 && month <= 12 {
		budget[month] = amount
		return month, nil //budget[month]
	}
	return 0, fmt.Errorf("invalid month values passed %v", 0)
}
func ExportExpenses() {}
