package main

import (
	"fmt"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func IsValidNumber(num string, base int) bool {
	for _, v := range num {
		if !strings.Contains(base36[:base], string(v)) {
			return false
		}
	}
	return true
}

func main() {
	var num string
	var base int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	fmt.Print("Enter base 1 - 36: ")
	fmt.Scan(&base)
	if base < 1 || base > 36 {
		fmt.Println("ERROR: wrong choise")
		return
	}
	if IsValidNumber(num, base) {
		fmt.Printf("Number %s is a correct representation of a number in base %v", num, base)
	} else {
		fmt.Printf("Number %s is NOT a correct representation of a number in base %v", num, base)
	}
}
