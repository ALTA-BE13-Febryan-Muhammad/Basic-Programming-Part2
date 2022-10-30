package main

import (
	"fmt"
)

func FaktorBilangan(n int) string {
	// your code here
	var number string
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			number = "faktorisasi"
		}
	}
	return number

}

func main() {
	var number int = 6
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilangan(number))

	var number2 int = 20
	fmt.Scanf("%d", &number2)
	fmt.Println(FaktorBilangan(number2))
}
