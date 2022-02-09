package main

import (
	"fmt"
	"strings"
)

const dictionary = "0123456789abcdefghijklmnopqrstuvwxyz" //hint

func isValidNumber(num string, base int) bool {
	for _, v := range num {
		if !strings.Contains(dictionary[:base], string(v)) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(" Program that validates if a string is a correct representation of a number in base B (`1 <= B <= 36`)")
	var num string
	var b int
	fmt.Println("Enter number:")
	fmt.Scan(&num)
	fmt.Println("Enter base:")
	fmt.Scan(&b)
	if isValidNumber(num, b) {
		fmt.Println("Number", num, "is in base", b)
	} else {
		fmt.Println("Number", num, "not in base", b)
	}
}
