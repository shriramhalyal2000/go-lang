package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
	fmt.Println("----newPattern---")
	// for m := 5; m >= 1; m++ {
	// 	for n := 1; n <= m; n++ {
	// 		fmt.Print(n)
	// 	}
	// 	fmt.Println()
	// }
}
