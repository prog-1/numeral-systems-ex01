package main

import (
	"fmt"
	"strings"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func isValidNumber(num string, base int) bool {
	if base < 1 || base > 36 {
		return false
	}
	if base >= 1 && base <= 36 {
		for _, v := range num {
			if !strings.Contains(base36[:base], string(v)) {
				return false
			}
		}
	}
	return true
}

func main() {
	var num string
	var base int
	fmt.Print("Enter number:")
	fmt.Scan(&num)
	fmt.Print("Enter base:")
	fmt.Scan(&base)
	if base >= 1 && base <= 36 {
		if isValidNumber(num, base) {
			fmt.Printf("String %v is a correct representation of a base %v", num, base)
		} else {
			fmt.Printf("String %v is incorrect representation of a base %v", num, base)
		}
	} else {
		fmt.Println("Base is incorrect")
	}
}
