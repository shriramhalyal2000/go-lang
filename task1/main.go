// creating a tax calculator with error handling cpabilities

package main

import (
	"errors"
	"fmt"
)

func getUsrInput(getText string) (float64, error) {
	var userInput float64
	fmt.Println(getText)
	fmt.Scanln(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Enter non zero positive number")
	}
	return userInput, nil
}

func calculateTax(revenue, expense, taxrate float64) (float64, float64, float64) {
	ebt := revenue - expense
	profit := ebt * (1 - taxrate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func main() {
	fmt.Println("Tax earnings, and Profit calculator")
	revenue, err := getUsrInput("Enter revenue:")
	if err != nil {
		fmt.Println(err)
		return
	}
	expense, err := getUsrInput("Enter expense:")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUsrInput("Enter tax rate:")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateTax(revenue, expense, taxRate)
	fmt.Printf("Earnings before tax is:%.2f\n", ebt)
	fmt.Printf("Profit on earnings after expense is:%.2f\n", profit)
	fmt.Printf("Earnings to profit ratio  is:%.2f\n", ratio)
}
