package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, fault := FileExists("readme")
	if fault != nil {
		return
	}
	fmt.Println(file)
	content, err := ReadFile("readme")
	if err != nil {
		return
	}
	fmt.Println(content)
}

// takes in file path as string, and returns contents and an potential error

func ReadFile(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) { // handles specific file not exist error
			fmt.Println("file not found!")
			return "", nil
		} else if errors.Is(err, os.ErrPermission) { // handles specific file permissione error
			fmt.Println("Error!, File permission error")
			return "", err
		}
		fmt.Println("Error in file handling")
		return "", err
	}
	fmt.Println("The file data is:")
	return string(data), nil

}

// takes in file as an argument and returns file exists or not as bool and a potential error

func FileExists(path string) (bool, error) {
	_, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) { // handles specific file not exist error
			fmt.Println("file not found!")
			return false, nil
		} else if errors.Is(err, os.ErrPermission) { // handles specific file permissione error
			fmt.Println("Error!Permision denied")
			return false, nil
		}
		fmt.Println("Error!, file not found.")
		return false, err
	}
	fmt.Println("File exists!")
	return true, nil
}
