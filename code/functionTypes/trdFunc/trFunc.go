package main

import "fmt"

func voting(age int, criminalHistory bool, nri bool) {
	if age < 18 {
		fmt.Println("Not eligible for voting")
	} else {
		if criminalHistory {
			fmt.Println("Not eligible for voting")
		} else if nri {
			fmt.Println("Not eligible for voting")
		} else {
			fmt.Println("Eligible for voting")
		}
	}
}

func main() {
	var age int
	var criminalHistory, nri bool
	fmt.Println("Enter your age: ", age)
	fmt.Scanln(&age)
	fmt.Println("Enter if your have criminal history or nri status:", criminalHistory, nri)
	fmt.Scanln(&criminalHistory, &nri)
	voting(age, criminalHistory, nri)
}
