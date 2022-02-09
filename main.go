package main

import (
	"fmt"
	"strings"
)

const dictionary = "0123456789abcdefghijklmnopqrstuvwxyz"

func isValidNumber(num string, base int) bool {
	if base >= 1 && base <= 36 {
		for _, v := range num {
			if !strings.Contains(dictionary[:base], string(v)) {
				return false
			}
		}
	}
	return true
}

func main() {
	var num string
	fmt.Println("Enter number: ")
	fmt.Scan(&num)
	var base int
	fmt.Println("Enter base: ")
	fmt.Scanln(&base)
	if isValidNumber(num, base) {
		fmt.Printf("String %s is a valid number for base %v", num, base)
	} else {
		fmt.Printf("String %s isn't a valid number for base %v", num, base)
	}
}
