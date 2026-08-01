package main

import (
	"fmt"

	bankops "example.comk/packgaes/bankOps"
	"example.comk/packgaes/comms"
)

func main() {
	var choice int
	comms.PresentOptions()
	fmt.Println("Enter your choice:")
	fmt.Scanln(&choice)
	userBalance := bankops.Operation(choice)
	fmt.Println("End balance is:", userBalance)
}
