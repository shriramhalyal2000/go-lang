package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("File handling")
	target_file := "readme"
	content, err := os.ReadFile(target_file)
	if err != nil {
		fmt.Println("Error! in finding file")
		return
	}
	fmt.Println("This reads readme file\n", string(content))
	fmt.Println("---------------------------------")
	filePath := "data.txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error!, file not found", filePath)
		return
	}
	fmt.Println("The file contents are:\n", string(data))
}
