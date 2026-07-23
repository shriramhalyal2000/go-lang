package main

import (
	"fmt"
	"math"
)

func adjustedInflationRate() {
	const inflationRate = 2.5
	var investmentAmount float64
	fmt.Println("Enter your invest ment under billion$:")
	fmt.Scanln(&investmentAmount)
	rateOfReturn := 5.5
	years := 10.0
	returnAmount := investmentAmount * math.Pow(1+rateOfReturn/100, years)
	inflationAdjustedAmt := returnAmount * math.Pow(1+inflationRate/100, years)
	fmt.Println("The investment amount adjusted to inflation is :", inflationAdjustedAmt)

}

func mini_invest() {
	invest, rate, time := 10000.0, 5.5, 10.0 // decalre multiple variables without explicit type or var def
	short_term_val := invest * math.Pow(1+rate/100, time)
	fmt.Println("short term return rate", short_term_val)

}

func investment() {
	var investment_amt float64 // mention var type only for undecalred vars
	return_rate := 5.5         // for decalred vals no need to decalre var type
	time := 10
	var future_val float64

	fmt.Print("Enter the invest amount:\n")
	fmt.Scanln(&investment_amt)

	future_val = investment_amt * math.Pow(1+return_rate/100, float64(time))
	fmt.Println("future adjusted value is ", future_val)

}

func main() {
	investment()
	mini_invest()
	adjustedInflationRate()
}
