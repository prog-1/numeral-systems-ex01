package main

import (
	"fmt"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func isValidNumber(num string, base int) bool {
	for _, v := range num {
		if !strings.Contains(base36[:base], string(v)) {
			return false
		}
	}
	return true
}

func main() {
	var number string
	var base int
	fmt.Print("Enter nummber: ")
	fmt.Scan(&number)
	fmt.Print("Enter base: ")
	fmt.Scan(&base)
	if isValidNumber(number, base) {
		fmt.Printf("Number %v is valid for base %v\n", number, base)
	} else {
		fmt.Printf("Number %v is not valid for base %v\n", number, base)
	}
}
