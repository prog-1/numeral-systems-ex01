package main

import (
	"fmt"
	"os"
)

const base36 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
	fmt.Println(`
	111101(2) -> 1 * 2^0 + 0 * 2^1 + 1 * 2^2 + 1 * 2^3 + 1 * 2^4 + 1 * 2^5 = 61(10)
	331(4) -> 1 * 4^0 + 3 * 4^1 + 3 *4^2 = 61(10)
	ABCD(16) -> 13 *16^0 + 12 * 16^1 + 11 * 16^2 + 10 * 16^3 = 43981(10)
	ABCD(20) -> 13 * 20^0 + 12 * 20^1 + 11 *20^2 + 10 20^3 = 84653 (10) `)
}
