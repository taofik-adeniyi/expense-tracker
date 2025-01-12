package main

import (
	"fmt"
	"os"
	"strconv"
)

func HandleAddition(f []string) {
	if len(f) != 6 {
		os.Exit(1)
	}
	var description string
	var amount int

	var second = f[2]
	var fourth = f[4]

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

	added := Add(description, amount)
	fmt.Printf("# Expense added successfully (ID: %v)\n", added)
}
