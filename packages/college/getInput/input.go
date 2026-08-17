package getinput

import "fmt"

func GetInputString(data string) string { //displays message to user and acceps string data and returns tring data
	var inputData string
	fmt.Println(data)
	fmt.Scanln(&inputData)
	return inputData
}

func GetInputFloat(data string) float64 { //displays message to user and accepts float data from user and returns it
	var inputFloat float64
	fmt.Println(data)
	fmt.Scanln(&inputFloat)
	return inputFloat
}
