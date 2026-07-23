package main

import "fmt"

var a int
var b int = 2
var c int = 4

func sub() {
	var g int
	var h int = 100
	x := 10
	g = h - x
	fmt.Println(g)
}

func add() {
	a = b + c
	fmt.Println(a)
}
func main() {
	fmt.Println("Variable decalred output")
	add()
	fmt.Println("Substraction function")
	sub()
}
