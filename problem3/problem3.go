package main

import (
	"fmt"
)

func PrimeNumber(number int) bool {
	// your code here
	var faktor int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			faktor++
		}
	}
	if faktor == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(PrimeNumber(3))
	fmt.Println(PrimeNumber(7))
	fmt.Println(PrimeNumber(10))
}
