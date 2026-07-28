package main

import "fmt"

func calculator(num1 float64, num2 float64) (float64, float64, float64, float64) {
	sum := num1 + num2
	sub := num1 - num2
	mul := num1 * num2
	div := num1 / num2

	return sum, sub, mul, div
}
func main() {
	var num1 float64
	var num2 float64
	fmt.Println("Enter two numbers :")
	fmt.Scanln(&num1, &num2)
	fmt.Println("The numbers with operated results are as follows")
	sum, sub, mul, div := calculator(num1, num2)
	fmt.Println("Sum of two numbers is : ", sum)
	fmt.Println("Substraction of two numbers is : ", sub)
	fmt.Println("Multiplication of two numbers is : ", mul)
	fmt.Printf("Division of two numbers is :  %.4f", div) // used Printf for value formatting and limiting divisionvalue after decimal to 4 digits
}
