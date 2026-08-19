package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	file := "readme.md"
	delete_file := "notreadme"
	filePath, err := CheckFile(file)
	if err != nil {
		return
	}
	fmt.Println("File status is:", filePath)
	fmt.Println(WriteToFile(file, []byte("This is written by function file write")))
	fileData, err := ReadFile(file)
	fmt.Println("File data is:", fileData)
	fmt.Println("Append data to file:", ApendData(file, []byte("\nAppended via function")))
	fmt.Println("Deleteing said file:", DeleteFlie(delete_file))
	size, modTime, err := FileStat(file)
	if err != nil {
		fmt.Println("Error! is:")
		return
	}
	fmt.Println("File size is:", size)
	fmt.Println("Mod time of file is:", modTime)
	fmt.Println("Copying file:", CopyFile(file, delete_file))
	defer fmt.Println("Deleting copied filed:", DeleteFlie(delete_file))
}

func CheckFile(path string) (bool, error) {
	_, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println("Error!, file doesnot exist")
			return false, err
		} else if errors.Is(err, os.ErrPermission) {
			fmt.Println("Error!, file permission error")
			return false, err
		}
		return false, err
	}
	return true, err

}
func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println("Error!, File doesnot exist")
			return "", err
		} else if errors.Is(err, os.ErrPermission) {
			fmt.Println("Error!, file permission error")
			return "", err
		}
		return "", err
	}
	fmt.Println("The file data is:")
	return string(data), nil

}

func WriteToFile(path string, data []byte) error { // overrites existing data
	return os.WriteFile(path, data, 0644)
}

// append data to a file
func ApendData(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644) // this step opens and appends file before writing new content
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println("Error!, writing file")
			return err
		}
		return err
	}
	defer file.Close()
	_, err = file.Write(data) // this step adds new data to open file
	return err
}

// deleteing a file
func DeleteFlie(path string) error {
	err := os.Remove(path)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Println("Error!, file not exist")
			return err
		} else if errors.Is(err, os.ErrPermission) {
			fmt.Print("Error!, file permission not exist")
			return err
		}
		return err
	}
	return nil
}

// get file statts

func FileStat(path string) (int64, time.Time, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 1, time.Time{}, err
	}
	size := info.Size()
	modTime := info.ModTime()
	return size, modTime, nil

}

// opy file form target to dstination

func CopyFile(srcPath, desPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(desPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}
	return dest.Sync()
}
