package main

import (
	"fmt"
	"strings"
)

const dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func isValidNumber(num string, base int) bool {
	for _, v := range num {
		if !strings.Contains(dictionary[:base], string(v)) {
			return false
		}
	}
	return true
}

func main() {
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
