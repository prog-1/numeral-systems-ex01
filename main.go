package main

import "fmt"

func main() {
	var base int
	var num string
	fmt.Println("Enter numbers")
	fmt.Scan(&num)
	fmt.Println("Enter base")
	fmt.Scan(&base)
	fmt.Println(isValidNumber())
	return
}

func isValidNumber(num string, base int) bool {
	for i := range num {
		if num[i] >= base {
			return false
		}
	}
	return true
}
