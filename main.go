package main

import (
	"fmt"
	"strings"
)

func isValidNumber(num string, base int) bool {
	base36 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if 1 <= base && base <= 36 {
		for _, v := range num {
			if !strings.Contains(base36[:base], string(v)) {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func main() {
	var num string
	fmt.Print("Enter number: ")
	fmt.Scan(&num)
	var base int
	fmt.Print("Enter base: ")
	fmt.Scan(&base)
	if 1 <= base && base <= 36 {
		fmt.Println(isValidNumber(num, base))
	} else {
		fmt.Println("ERROR: Incorrect base")
	}
}
