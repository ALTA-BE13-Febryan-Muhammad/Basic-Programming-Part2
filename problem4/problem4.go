package main

import "fmt"

func Palindrome(input string) bool {
	// your code here
	reversedStr := ""
	for i := len(input) - 1; i >= 0; i-- {
		reversedStr += string(input[i])
	}
	for i := range input {
		if input[i] != reversedStr[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Palindrome("civic"))       // true
	fmt.Println(Palindrome("katak"))       // true
	fmt.Println(Palindrome("kasur rusak")) // true
	fmt.Println(Palindrome("kupu-kupu"))   // false
	fmt.Println(Palindrome("lion"))        // false
}
