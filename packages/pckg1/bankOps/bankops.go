package bankops

import (
	"fmt"

	"example.comk/packgaes/fileops"
	inputfunc "example.comk/packgaes/inputFunc"
)

const accounBalanceFile = "balance.txt"

func Operation(choice int) float64 { // bank operation logic to get balance, withdraw money and dipost money
	balance, err := fileops.GetFloatFromFile(accounBalanceFile)
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
			fileops.WriteToFile(balance, accounBalanceFile)
		} else if choice == 2 {
			deposit = inputfunc.GetInput("Enter deposite amount:")
			balance += deposit
			fmt.Println("The updated balance is: ", balance)
			fileops.WriteToFile(balance, accounBalanceFile)
		} else if choice == 3 {
			if balance >= 100 {
				withdraw := inputfunc.GetInput("Withdraw amount:")
				if withdraw >= 100 {
					balance -= withdraw
					fmt.Println("Updated balance after withdraw amount is :", balance)
				} else if withdraw < 100 {
					fmt.Println("Withdraw amount must be greater than 100")
				} else {
					fmt.Println("Withdraw amount should alwasy be a positive number")
				}
				fileops.WriteToFile(balance, accounBalanceFile)
			}
		} else {
			fmt.Println("Invalid choice of entry:", choice)
		}
		fmt.Println("\nAvailable choice are: 1.Balance, 2. Deposite, 3. Withdraw, 4. Exit")
		fmt.Scanln(&choice)
	}
	return balance

}
