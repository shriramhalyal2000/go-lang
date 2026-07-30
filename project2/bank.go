package main

import "fmt"

func getUserInput(getText string) float64 {
	var usrInput float64
	fmt.Println(getText)
	fmt.Scanln(&usrInput)
	return usrInput
}

func bankTransaction(choice int) float64 {
	balance := 999999.9
	var withdraw float64
	var deposit float64
	switch choice {
	case 1:
		fmt.Println("Available initial balance is:", balance)
	case 2:
		balance += deposit
		fmt.Println("Updated balance after deposit is:", balance)
	case 3:
		balance -= withdraw
		fmt.Println("Updated balance after withdraw is:", balance)
	case 4:
		break
	default:
		fmt.Println("Invalid choice try again")
		choice = int(getUserInput("Please enter again from available menue, 1. Display balance, 2.Deposit amount, 3. Withdraw amount, 4. Exit"))
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
	userBalance := bankTransaction(choice)
	fmt.Println("End balance is:", userBalance)

}
