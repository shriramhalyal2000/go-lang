package main

import (
	"errors"
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
	balanceTxt := fmt.Sprint(balance) // formats balance, Formatting string from stored value and passing it in to a var
	os.WriteFile(accounBalanceFile, []byte(balanceTxt), 0644)
}
func fetchBalance() (float64, error) {
	data, err := os.ReadFile(accounBalanceFile) // var to hold file const, looks for the file to read, but also returns error if it doesnot find any(err), points to file from that function
	if err != nil {
		return 0, errors.New("failed to find balace repot file") // returns error message when balance file not found, stopps code from crshing.
	}
	balanceText := string(data)                         //convert data in to string, balance is float64 --> pass file into variable to convert its data into string
	balance, err := strconv.ParseFloat(balanceText, 64) //convert data back to encoded float64 --> convert the string data to float
	if err != nil {
		return 0, errors.New("Failed to parse the file data")
	}
	return balance, nil
}
func operation(choice int) float64 { // bank operation logic to get balance, withdraw money and dipost money
	balance, err := fetchBalance()
	if err != nil {
		fmt.Println("---error---")
		fmt.Println(err) //prints user defined error here
		fmt.Println("---error---")
	}
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
