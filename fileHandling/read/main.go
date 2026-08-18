package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := ReadFile("readme")
	if err != nil {
		return
	}
	fmt.Println(content)
}
func ReadFile(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error in file handling")
		return "", err
	}
	fmt.Println("The file data is:")
	return string(data), nil

}
