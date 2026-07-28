package main

import "fmt"

func profitCalculator() {
	var revenue float64
	var expense float64
	var taxRate float64
	fmt.Println("Enter your revenue per year:")
	fmt.Scanln(&revenue)
	fmt.Println("Enter your expense of businees per year:")
	fmt.Scanln(&expense)
	fmt.Println("Enter the tax rate of your state:")
	fmt.Scanln(&taxRate)

	earningBeforeTax := revenue - expense //pure earnings after deducting expense for yearly earning to get profit
	fmt.Println("Earnings before tax is revenue-expense :", earningBeforeTax)

	profit := earningBeforeTax * (1 - taxRate/100)
	fmt.Println("Profit after applying tas is:", profit) // pure profit after tax and expense deducations

	earningsAfterTax := float64(earningBeforeTax * (1 - taxRate/100)) // earnings before tax is matched with taxRate on that earnings as whole
	fmt.Println("Earnings after tas is :", earningsAfterTax)

	taxPaidIs := float64(earningBeforeTax - earningsAfterTax)
	fmt.Printf("Total tax paid is :%v", taxPaidIs)

}

func main() {
	fmt.Println("This is a profit calculator:")
	profitCalculator()
}
