package main

import "fmt"

// calculateTax computes tax based on progressive slabs under the New Tax Regime.
func calculateTax(earnings float64) (float64, float64, float64) {
	// Standard Deduction for salaried employees
	standardDeduction := 75000.0
	taxableIncome := earnings - standardDeduction
	if taxableIncome < 0 {
		taxableIncome = 0
	}

	var grossTax float64

	// Slab calculations
	if taxableIncome > 1200000 {
		// Income above 12L taxed at 15% (for illustration if earnings > 12L)
		grossTax += (taxableIncome - 1200000) * 0.15
		grossTax += (1200000 - 800000) * 0.10
		grossTax += (800000 - 400000) * 0.05
	} else if taxableIncome > 800000 {
		// 10% on amount between 8L and 12L
		grossTax += (taxableIncome - 800000) * 0.10
		grossTax += (800000 - 400000) * 0.05
	} else if taxableIncome > 400000 {
		// 5% on amount between 4L and 8L
		grossTax += (taxableIncome - 400000) * 0.05
	}

	// Section 87A Rebate: Full tax rebate if taxable income is <= 12,00,000
	var finalTax float64
	if taxableIncome <= 1200000 {
		finalTax = 0.0
	} else {
		// Add 4% Health & Education Cess if tax is due
		finalTax = grossTax * 1.04
	}

	return taxableIncome, grossTax, finalTax
}

func main() {
	var earnings float64

	fmt.Print("Enter Earnings Before Tax: ")
	fmt.Scanln(&earnings)

	taxableIncome, grossTax, finalTax := calculateTax(earnings)

	fmt.Println("\n--- Tax Breakdown ---")
	fmt.Printf("Gross Earnings:      ₹%.2f\n", earnings)
	fmt.Printf("Taxable Income:      ₹%.2f (after ₹75,000 Std. Deduction)\n", taxableIncome)
	fmt.Printf("Gross Tax (Slabs):   ₹%.2f\n", grossTax)
	fmt.Printf("Final Tax Payable:   ₹%.2f (after Sec 87A rebate & Cess)\n", finalTax)
}
