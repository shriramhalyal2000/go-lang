package main

import (
	"fmt"
	"os"

	getinput "example.com/college/getInput"
)

const dataFile = "data.txt"

func WriteStringToFile(data, filename string) { // doesnot return anything only takes data as arguments
	dataTxt := fmt.Sprint(data)
	os.WriteFile(filename, []byte(dataTxt), 0644)
}
func WriteFloatToString(value float64, filename string) { //writes float data to file, as input into txt file
	valueTxt := fmt.Sprintln(value)
	os.WriteFile(filename, []byte(valueTxt), 0644)
}

func AppendStringToFile(data, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
		return err
	}
	defer file.Close()
	return err
}

func main() {
	fmt.Println("This is voting right eligibility check:")
	firstName := getinput.GetInputString("Enter your name:")
	WriteStringToFile(firstName, dataFile)
	lastName := getinput.GetInputString("Enter your last name:")
	age := getinput.GetInputFloat("Enter your age:")

	if age >= 18 {
		crimeHistory := getinput.GetInputString("Do you have criminal histroy? y/n:")
		resident := getinput.GetInputString("Are you currently living in India? y/n:")
		if crimeHistory == "y" && resident == "n" {
			fmt.Println("You cant vote", firstName, lastName)
		} else {
			fmt.Println("You are eligible for vaoting", firstName, lastName)
		}
	} else {
		fmt.Println("you are not eligible for voting.")
	}
}
