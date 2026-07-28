package main

import "fmt"

func myTax(earnings float64) (float64, float64, float64) {
	standardDeduction := 75000.0
	taxableIncome := earnings - standardDeduction
	if taxableIncome < 40000 {
		taxableIncome = 0.0
	}
	var grossTax float64
	if taxableIncome > 1200000 {
		grossTax += (taxableIncome - 1200000) * 0.15
		grossTax += (1200000 - 800000) * .10
		grossTax += (800000 - 400000) * .005
	} else if taxableIncome > 800000 {
		grossTax += (taxableIncome - 800000) * .01
		grossTax += (800000 - 400000) * .005
	} else if taxableIncome > 400000 {
		grossTax += (taxableIncome - 400000) * .05
	}
	var finalIncomeTax float64
	if taxableIncome <= 1200000 {
		finalIncomeTax = 0.0
	} else {
		finalIncomeTax = grossTax * 1.04
	}
	return taxableIncome, grossTax, finalIncomeTax
}

func main() {
	var earnings float64
	fmt.Println("Enter your income per year:")
	fmt.Scanln(&earnings)
	taxableIncome, grossTax, finalIncomeTax := myTax(earnings)
	fmt.Println("--Tax Calculations--")
	fmt.Printf("Taxble income is : %.2f\n", taxableIncome)
	fmt.Printf("Final Gross tax is : %.2f\n", grossTax)
	fmt.Printf("Final income tax on your taxable income is : %.2f\n", finalIncomeTax)

}
