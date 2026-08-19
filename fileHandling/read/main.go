package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	path := "readme1"
	//path1 := "readmenot"
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
	fmt.Println("create file:", CreateFile(path))
	fmt.Println("File created if not already:", CreteFileIfNotExist(path))
	fmt.Println("Write to file with data:", WriteToFile("readme1.txt", []byte("The file is been overwritten with this function")))
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

// create file function

func CreateFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil // if function suceeds no error
}

// create file if not created

func CreteFileIfNotExist(path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_EXCL|os.O_CREATE, 0666) // checks path for file, read write perms exclusively, and creates if doesnot exist.
	// if the file exists thrpws error
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// write to file, which overrides exiting content with new one

func WriteToFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
