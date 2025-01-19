package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"
)

type Expense struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Category    string    `json:"category"`
}

var defaultDB = make([]Expense, 0)

var monthsOfYear = map[int]string{
	1:  "january",
	2:  "february",
	3:  "march",
	4:  "april",
	5:  "may",
	6:  "june",
	7:  "july",
	8:  "august",
	9:  "september",
	10: "october",
	11: "november",
	12: "december",
}

func Add(description string, amount int, month int) (int, error) {
	fmt.Printf("expense-tracker: description: %v amount: %v month: %v\n", description, amount, month)

	// check budget balance
	b, err := getBudget(month)

	// check if balance for month has a budget
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("the initial balance: %v \n", b)

	// checking budget
	// if amount is greater than budget balance return error an terminate function execution

	valid, err := budgetExceeded(amount, month)
	if err != nil {
		fmt.Println("Error:", err)
		// return 0, err
	}
	if valid {
		fmt.Printf("unable to add expense of amount: %v for: %v excedeed budget: %v\n", amount, description, b)
		fmt.Println("")
		fmt.Print("Set up budget based on month given january is 1 febuary is 2")
		fmt.Scanln(&month)
		fmt.Print("budget amount")
		fmt.Scanln(&amount)
	}
	var amountToSet = b - amount
	budget, err := SetBudget(month, amountToSet)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
	}
	fmt.Printf("the budget saved: %v\n", budget)

	// fetch expense file
	fileByte, err := getFileContent(dbFileName)
	fmt.Println("fileByte", fileByte)
	if err != nil {
		fmt.Println("add get file err")
		fmt.Printf("Error: %v", err.Error())
		return 0, err
	}
	var data []Expense
	if len(fileByte) == 0 {
		data = []Expense{}
	} else {
		err = json.Unmarshal(fileByte, &data)
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}
	fmt.Println("data", data)

	var newExpense = Expense{
		Id:          len(defaultDB) + 1,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
		Category:    "default category",
	}
	data = append(data, newExpense)
	fmt.Println("data:after", data)

	savedByte, err := json.Marshal(&data)
	fmt.Println("Byte to save", savedByte)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	err = saveToFileDb(dbFileName, savedByte)
	if err != nil {
		fmt.Println("Error saving file:", err.Error())
	}

	// read the expense file
	// create a new expense and add to file
	// return the created expense

	return 0, nil
}

func List() []Expense {
	// 	# ID  Date       Description  Amount
	// # 1   2024-08-06  Lunch        $20
	// # 2   2024-08-06  Dinner       $10
	return []Expense{}
}

func (e Expense) Summary(month *int) {
	// # Total expenses: $30
	fmt.Printf("Total expenses: $%v", e.Amount)
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
	// var budget map[int]int = make(map[int]int, 0) // map month to the months budget
	// getBudget()

	budget := getAppBudget()

	if month >= 1 && month <= 12 {
		budget[month] = amount

		fmt.Printf("your budget for the month of: %v is: %v\n", monthsOfYear[month], budget)
		budgetByte, err := json.Marshal(budget)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		saveToFileDb(balanceFileName, budgetByte)

		return month, nil //budget[month]
	}
	return 0, fmt.Errorf("invalid month values passed %v", 0)
}

func SetDefaultBalance() {

	content, err := getFileContent(balanceFileName)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		os.Exit(1)
	}
	if len(content) > 0 {
		// fmt.Println("content exists already for balance database")
		return
	}
	var budgets = map[int]int{
		1:  0,
		2:  0,
		3:  0,
		4:  0,
		5:  0,
		6:  0,
		7:  0,
		8:  0,
		9:  0,
		10: 0,
		11: 0,
		12: 0,
	}

	budgetBytes, err := json.Marshal(budgets)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		os.Exit(1)
	}
	err = saveToFileDb(balanceFileName, budgetBytes)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		os.Exit(1)
	}
}
func getBudget(month int) (int, error) {

	fileByteContent, err := getFileContent(balanceFileName)
	var budget map[int]int

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		// os.Exit(1)
	}

	err = json.Unmarshal(fileByteContent, &budget)
	if err != nil {
		fmt.Printf("Error: %v \n", err.Error())
	}

	if val, ok := budget[month]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("no budget for the month of %v", month)
}
func getAppBudget() map[int]int {
	fileByteContent, err := getFileContent(balanceFileName)
	var budget map[int]int

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}
	err = json.Unmarshal(fileByteContent, &budget)
	if err != nil {
		fmt.Printf("Error: %v \n", err.Error())
	}

	return budget
}
func budgetExceeded(amount int, month int) (bool, error) {
	budget, err := getBudget(month)
	if err != nil {
		return true, err
	}

	if budget > amount {
		return false, nil
	}
	return true, nil
}
func ExportExpenses() {}

func saveToFileDb(fileName string, content []byte) error {
	var filePath = fileName

	err := os.WriteFile(filePath, content, 0644)
	if err != nil {
		fmt.Printf("Error saving to file storage\n")
		os.Exit(1)
		return fmt.Errorf("Error saving to file storage: %v \n", err)
	}
	fmt.Println("file content saved succesfully")
	return nil
}

func getFileContent(fileName string) ([]byte, error) {
	fsys := os.DirFS(".")

	bytes, err := fs.ReadFile(fsys, fileName)

	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func createFileIfNotExists(fileName string) {

	tF, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		fmt.Println("Creating db file.....")
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Error: %v", err.Error())
		}
		fmt.Printf("Db file created: %v\n", file.Name())
	}
	defer tF.Close()

}
