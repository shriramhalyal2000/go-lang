package main

import "fmt"

var a int
var b int = 2
var c int = 4

func mul() {
	var num1 int
	var num2 int
	var num3 int

	fmt.Println("Enter num for multiplication:")
	fmt.Scanln(&num2, &num3)
	num1 = num2 * num3
	fmt.Print(num1)
}

func sub() {
	var g int
	var h int = 100
	x := 10
	g = h - x
	fmt.Println("Multiplied number is:\n", g)
}

func add() {
	a = b + c
	fmt.Println(a)
}
func div() {
	var d float64
	var y float64
	var u float64

	fmt.Println("enter your div numbers:")
	fmt.Scanln(&y, &u)
	d = y / u
	fmt.Println("divided number is : \n", d)
}
func main() {
	fmt.Println("Variable decalred output")
	add()
	fmt.Println("Substraction function")
	sub()
	mul()
	fmt.Print("\n")
	div()
}
