package main

import (
	"fmt"
	"os"
	"strconv"
)

func HandleAddition(f []string) {

	if len(f) != 8 {
		os.Exit(1)
	}
	var description string
	var amount int
	var month int

	var second = f[2]
	var fourth = f[4]
	var sixth = f[6]
	fmt.Println("here....", sixth)

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
