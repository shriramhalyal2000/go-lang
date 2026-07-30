package main

import (
	"errors"
	"fmt"
	"os"
)

const userDataFile = "data.txt"

func getUserInput(getText string) string { //custom user iput function
	var usrInput string
	fmt.Println(getText)
	fmt.Scanln(&usrInput)
	return usrInput
}

func storeDataToFile(data string) { // file write functiont o store user data
	dataToWrite := fmt.Sprint(data)                       // string formatting data into a variable
	os.WriteFile(userDataFile, []byte(dataToWrite), 0644) // write data into textfile with file perms
}

func fetchUsrData() (string, error) {
	data, err := os.ReadFile(userDataFile) // fetchs and reads file
	if err != nil {
		return "", errors.New("File doesnot exist")
	}
	return string(data), nil
}

func askUser(userData string) string { //asks user to write some string data
	userData = getUserInput("Enter any string data")
	storeDataToFile(userData)
	return userData
}

func main() {
	writeSomething := askUser("Write Something without using space, but use -, _")
	fmt.Println("Successfully saved", writeSomething)
	showSomething, err := fetchUsrData()
	if err != nil {
		fmt.Println("---Error---", err)
	}
	fmt.Println("Fetched data from file", showSomething)
}
