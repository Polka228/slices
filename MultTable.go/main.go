package main

import "fmt"

func MultTable() (s [10][10]int) {
	for i := range s {
		for j := range s[i] {
			s[i][j] = i * j
		}
	}
	// s[0][1] = 0*1
	return s

}

func main() {
	fmt.Println(MultTable())

}
