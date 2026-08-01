package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteToFile(value float64, fileName string) { //writes balance to external file
	valueTxt := fmt.Sprint(value) // formats balance, Formatting string from stored value and passing it in to a var
	os.WriteFile(fileName, []byte(valueTxt), 0644)
}
func GetFloatFromFile(fileName string) (float64, error) { //recives file name as input returns data
	data, err := os.ReadFile(fileName) // var to hold file const, looks for the file to read, but also returns error if it doesnot find any(err), points to file from that function
	if err != nil {
		return 0, errors.New("failed to find file") // returns error message when balance file not found, stopps code from crshing.
	}
	valueText := string(data)                       //convert data in to string, balance is float64 --> pass file into variable to convert its data into string
	value, err := strconv.ParseFloat(valueText, 64) //convert data back to encoded float64 --> convert the string data to float
	if err != nil {
		return 0, errors.New("Failed to parse the file data")
	}
	return value, nil
}
