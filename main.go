package main

import (
	"fmt"
	"os"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ" //haha

func isValidNumber(num string, base int) bool {
	for _, r := range num {
		if byte(r) > base36[base] || base == 0 || base == 1 { // i found only theories about base 0 or 1, but noone use them in any practical case, so i think false here will be better
			return false
		}
	}
	return true
}
func emtyMain() {
	var num string
	var base int
	fmt.Print("Enter number you want to check: ")
	fmt.Scan(&num)
	fmt.Print("Enter base: ")
	fmt.Scan(&base)
	if base < 0 { // this case is pretty interesting
		fmt.Println(" You are trying to check number with non-standard positional numeral system. This program cannot do this at the moment. Please check wikipedia https://en.wikipedia.org/wiki/Negative_base and try to check number on your own.")
		os.Exit(0)
	}
	if isValidNumber(num, base) {
		fmt.Printf("Number %s is valid for base %v", num, base)
	} else {
		fmt.Printf("! Number %s is invalid for base %v! ", num, base)
	}
}

func main() {
	emtyMain()
}
