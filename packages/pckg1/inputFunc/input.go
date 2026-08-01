package inputfunc

import "fmt"

func GetInput(getText string) float64 { // displays message to user to get input from
	var userInput float64
	fmt.Println(getText)
	fmt.Scanln(&userInput)
	return userInput
}
