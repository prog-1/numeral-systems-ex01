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
	var num string
	fmt.Println("Enter number: ")
	fmt.Scan(&num)
	var base int
	fmt.Println("Enter base: ")
	fmt.Scan(&base)
	if base < 1 {
		fmt.Println("ERR: non-existent base")
	} else if base > 36 {
		fmt.Println("The program does not support this base.")
	} else {
		if isValidNumber(num, base) {
			fmt.Printf("%s is a valid number for base %v", num, base)
		} else {
			fmt.Printf("%s is not a valid number for base %v", num, base)
		}
	}
}
