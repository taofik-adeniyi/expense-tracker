package main

import (
	"encoding/csv"
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

type Expenses []Expense

func (e Expenses) formatExpenses() {
	// Print the header
	fmt.Printf("%-4v %-4s %-12s %-15s %6s\n", "#", "ID", "Date", "Description", "Amount")

	for key, value := range e {
		fmt.Printf("%-4v %-4d %-12s %-15s $%d\n", "#", key+1, value.Date.Format("2006-01-02"), value.Description, value.Amount)
	}

}

// var defaultDB = make([]Expense, 0)

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

func Summary() (Expenses, error) {
	content, err := getFileContent(dbFileName)
	if err != nil {
		return []Expense{}, fmt.Errorf("error: %v", err)
	}
	if len(content) == 0 {
		return []Expense{}, fmt.Errorf("no expense to be displayed")
	}
	var lists []Expense
	err = json.Unmarshal(content, &lists)
	if err != nil {
		return []Expense{}, fmt.Errorf("error: %v", err)
	}
	return lists, nil
}

func ListExpenses() (Expenses, error) {
	content, err := getFileContent(dbFileName)
	if err != nil {
		return []Expense{}, fmt.Errorf("error: %v", err)
	}
	if len(content) == 0 {
		return []Expense{}, fmt.Errorf("no expense to be displayed")
	}
	var lists []Expense
	err = json.Unmarshal(content, &lists)
	if err != nil {
		return []Expense{}, fmt.Errorf("error: %v", err)
	}
	return lists, nil
}
func Add(description string, amount int, month ...int) (int, error) {
	var monthValue int
	if len(month) > 0 {
		monthValue = month[0]
	} else {
		monthValue = int(time.Now().Month())
	}
	var budget int
	b, err := getBudget(monthValue)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	budget = b

	if b == 0 {
		// request user to set a budget before he can continue
		fmt.Printf("You do not have a budget, setup budget for: %v\n", monthsOfYear[monthValue])
		fmt.Println("Set up a budget:")
		var budgit int
		fmt.Scanln(&budgit)
		res, err := SetBudget(monthValue, budgit)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		budget = res
	}

	valid, err := budgetExceeded(amount, monthValue)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if valid {
		fmt.Printf("unable to add expense of amount: %v for expense: %v excedeed budget: %v\n", amount, description, budget)
		fmt.Println("")
		fmt.Println("Enter number between 1-12 1 means January and 12 is December")
		fmt.Scanln(&monthValue)
		if monthValue < 1 && monthValue > 12 {
			fmt.Println("The month value has to be in range 1-12 not more not less")
			fmt.Println("Enter number between 1-12 1 means January and 12 is December")
			fmt.Scanln(&monthValue)
		}
		fmt.Printf("Enter Budget Amount\n")
		fmt.Scanln(&amount)
	}
	var amountToSet = budget - amount
	_, err = SetBudget(monthValue, amountToSet)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
	}
	fileByte, err := getFileContent(dbFileName)
	if err != nil {
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

	var newExpense = Expense{
		Id:          len(data) + 1,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
		Category:    "default_category",
	}
	fmt.Println("b:amount", amount)
	data = append(data, newExpense)
	fmt.Println("a:amount", amount)

	savedByte, err := json.Marshal(&data)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	err = saveToFileDb(dbFileName, savedByte)
	if err != nil {
		fmt.Println("Error saving file:", err.Error())
	}
	return newExpense.Id, nil
}

func (e Expenses) Summary(month ...int) string {
	var total int
	var responseString string
	if len(month) > 0 && month[0] >= 1 && month[0] <= 12 {
		for _, value := range e {
			if value.Date.Month() == time.Month(month[0]) {
				total += value.Amount
			}
		}
		responseString = fmt.Sprintf("Total expenses for %v: $%d", time.Month(month[0]), total)
	} else {
		for _, value := range e {
			total += value.Amount
		}
		responseString = fmt.Sprintf("Total expenses: $%d", total)
	}
	return responseString
}

func (e Expenses) SummaryByCategory(category string) string {
	var total int
	var responseString string
	var found bool

	for _, value := range e {
		if value.Category == category {
			found = true
			total += value.Amount
		}
	}
	responseString = fmt.Sprintf("Total expenses: $%d", total)
	if !found {
		return fmt.Sprintf("No expense for category: %s\n", category)
	}
	return responseString
}

func Delete(id int) (int, error) {
	var data Expenses
	var newData Expenses
	contentByte, err := getFileContent(dbFileName)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(contentByte, &data)
	if err != nil {
		return 0, err
	}
	for _, value := range data {
		if value.Id != id {
			newData = append(newData, value)
		}
	}
	toSaveByte, err := json.Marshal(newData)
	if err != nil {
		return 0, nil
	}
	err = saveToFileDb(dbFileName, toSaveByte)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func Update(id int, description string, amount int, category string) (string, error) {

	var contentData Expenses
	contentByte, err := getFileContent(dbFileName)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(contentByte, &contentData)
	if err != nil {
		return "", err
	}
	var found bool
	for key, value := range contentData {
		if value.Id == id {
			found = true
			if category != "" {
				contentData[key].Category = category
			}
			if amount != 0 {
				contentData[key].Amount = amount
			}
			if description != "" {
				contentData[key].Description = description
			}
			contentData[key].Date = time.Now()
		}
	}
	if !found {
		return "", fmt.Errorf("Expense with ID %d not found", id)
	}

	contentByteToSave, err := json.Marshal(contentData)
	if err != nil {
		return "", err
	}
	err = saveToFileDb(dbFileName, contentByteToSave)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Expense updated successfully (ID: %d)", id), nil

}
func SetBudget(month int, amount int) (int, error) {
	budget := getAppBudget()

	if month >= 1 && month <= 12 {
		budget[month] = amount

		budgetByte, err := json.Marshal(budget)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		saveToFileDb(balanceFileName, budgetByte)

		return amount, nil
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

func createFileIfNotExists(fileName string) *os.File {

	tF, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Error file does not exists: %v\n", err)
		fmt.Println("Creating file .....")
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Error: %v", err.Error())
		}
		fmt.Printf("File created: %v\n", file.Name())
	}
	defer tF.Close()
	return tF

}

func transformObjectsToArrays(objects []map[string]interface{}, keys []string) [][]interface{} {
	var result [][]interface{}

	for _, obj := range objects {
		var arr []interface{}

		for _, key := range keys {
			if key == "amount" {
				value := fmt.Sprintf("$%v", obj[key])
				arr = append(arr, value) // Extract values in the defined key order
			} else if key == "date" {
				str, ok := obj[key].(string)
				if !ok {
					log.Fatal("value is not a string")
				}
				parsedTime, err := time.Parse(time.RFC3339Nano, str)

				if err != nil {
					log.Fatal("cant convert to date ")
				}
				formattedTime := parsedTime.Format("2006-01-02 15:04:05")
				arr = append(arr, formattedTime)

			} else {
				arr = append(arr, obj[key]) // Extract values in the defined key order
			}

		}

		result = append(result, arr)
	}

	return result
}

func csvExport(fileName string, data [][]interface{}) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var headers = []string{"id", "description", "category", "date", "amount"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("error writing to csv: %v", err)
	}

	for _, value := range data {
		strRow := make([]string, len(value))
		for i, v := range value {
			strRow[i] = fmt.Sprintf("%v", v) // Convert to string using fmt.Sprintf
		}
		if err := writer.Write(strRow); err != nil {
			return fmt.Errorf("error writing to csv: %v", err)
		}
	}
	return nil
}
func ExportExpensesToCsv(fileName string) {
	// var content Expenses
	var newContent []map[string]interface{}
	fileBytes, err := getFileContent(dbFileName)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// err = json.Unmarshal(fileBytes, &content)
	err = json.Unmarshal(fileBytes, &newContent)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// Define the desired key order
	keys := []string{"id", "description", "category", "date", "amount"}
	dttt := transformObjectsToArrays(newContent, keys)
	err = csvExport(fileName, dttt)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("Data exported and CSV File created: %v", fileName)
}
