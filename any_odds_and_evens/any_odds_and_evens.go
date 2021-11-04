package main

import "fmt"

func main() {
	fmt.Println("Enter numbers: ")
	var s []int
	fmt.Scan(&s)
	if AnyIsOdd() == true{
		fmt.Println("all numbers are negative")
	} else {
		fmt.Println("There is imposter")
	}
	fmt.Println("Enter numbers: ")
	var b [4]int
	fmt.Scan(&b)
	if AnyIsEven() == true{
		fmt.Println("all numbers are positive")
	} else {
		fmt.Println("There is negative number")
	}

	return
}

func AnyIsOdd(s []int) bool {
	for _, v := range s {
		if v%2 == 1 {
			return true
		}
		return false
	}

func AnyIsEven(s []int) bool{
	for _, v := range s {
		if v%2 == 0 {
			return true
		}
	}
	return false
}
