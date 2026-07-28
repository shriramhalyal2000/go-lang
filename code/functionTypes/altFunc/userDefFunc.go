package main

import (
	"fmt"
)

func getUserInput(getText string) float64 {
	var inputText float64
	fmt.Println(getText)
	fmt.Scanln(&inputText)
	return inputText
}

func calculateFinancials(revenue, expense, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
func main() {
	revenue := getUserInput("Revenue:")
	expenses := getUserInput("Expenses:")
	taxRate := getUserInput("TaxRate:")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Printf("Earnings before tax:\n%.4f", ebt)
	fmt.Printf("\nProfit after tax:%.4f\n", profit)
	fmt.Printf("Ratio:%.4f", ratio)
}
