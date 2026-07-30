package main

import (
	"fmt"
)

func getInput(getText string) float64 {
	var userInput float64
	fmt.Println(getText)
	fmt.Scanln(&userInput)
	return userInput
}

func operation(choice int) float64 {
	balance := 999999.0
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
		} else if choice == 2 {
			deposit = getInput("Enter deposite amount:")
			balance += deposit
			fmt.Println("The updated balance is: ", balance)
		} else if choice == 3 {
			if balance >= 100 {
				withdraw := getInput("Withdraw amount:")
				balance -= withdraw
				fmt.Println("Updated balance after withdraw amount is :", balance)
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
