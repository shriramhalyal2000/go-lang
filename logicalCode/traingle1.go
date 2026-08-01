package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
	fmt.Println("----newPattern----")
	fmt.Println("----Revtriangle----")
	for m := 5; m >= 1; m-- {
		for n := 1; n <= m; n++ {
			fmt.Print(n)
		}
		fmt.Println()
	}
	fmt.Println("----newPattern----")
	for k := 10; k >= 1; k-- {
		for l := 1; l <= k; l++ {
			fmt.Print(l)
		}
		fmt.Println()
	}
	fmt.Println("----newPattern----")
	for i := 5; i >= 1; i-- { //rows
		for j := i; j >= 1; j-- { //columns
			fmt.Print(j)
		}
		fmt.Println()
	}
	fmt.Println("----newPattern----")
	for i := 0; i >= 5; i-- {
		for k := 0; k < 5; k++ {
			fmt.Print(k)
		}
		for j := 0; j < 5-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
