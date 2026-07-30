package main

import (
	"fmt"
	"os"
	"strconv"
)

const accounBalanceFile = "balance.txt"

func getInput(getText string) float64 { // displays message to user to get input from
	var userInput float64
	fmt.Println(getText)
	fmt.Scanln(&userInput)
	return userInput
}

func toFile(balance float64) { //writes balance to external file
	balanceTxt := fmt.Sprint(balance) // formats balance
	os.WriteFile(accounBalanceFile, []byte(balanceTxt), 0644)
}
func fetchBalance() float64 {
	data, _ := os.ReadFile(accounBalanceFile)         // var to hold file const
	balanceText := string(data)                       //convert data in to string, balance is float64
	balance, _ := strconv.ParseFloat(balanceText, 64) //convert data back to encoded float64
	return balance
}
func operation(choice int) float64 { // bank operation logic to get balance, withdraw money and dipost money
	balance := fetchBalance()
	var deposit float64
	for choice != 4 {
		if choice == 1 {
			balance += deposit
			if balance > 0 {
				fmt.Println("The available balance is :", balance)
				//break // breaks out of loop completely
			} else {
				fmt.Println("You dont have displayable balance")
			}
			toFile(balance)
		} else if choice == 2 {
			deposit = getInput("Enter deposite amount:")
			balance += deposit
			fmt.Println("The updated balance is: ", balance)
			toFile(balance)
		} else if choice == 3 {
			if balance >= 100 {
				withdraw := getInput("Withdraw amount:")
				if withdraw >= 0 {
					balance -= withdraw
					fmt.Println("Updated balance after withdraw amount is :", balance)
				} else {
					fmt.Println("Withdraw amount should alwasy be a positive number")
				}
				toFile(balance)
			}
		} else {
			fmt.Println("Invalid choice of entry:", choice)
		}
		fmt.Println("\nAvailable choice are: 1.Balance, 2. Deposite, 3. Withdraw, 4. Exit")
		fmt.Scanln(&choice)
	}
	return balance

}

func main() {
	fmt.Println("Welcome to Go Bank app")
	fmt.Println("Available choices:")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposite Money")
	fmt.Println("3. Withdraw Money:")
	fmt.Println("4. Exit")
	var choice int
	fmt.Println("Enter your choice:")
	fmt.Scanln(&choice)
	userBalance := operation(choice)
	fmt.Println("End balance is:", userBalance)
}
