package main

import (
	"bufio"
	"fmt"
	"os"
)

func isNumber(num string) bool {
	for i := range num {
		if num[i] < 48 || num[i] > 90 || (num[i] > 57 && num[i] < 65) {
			return false
		}
	}
	return true
}

func isValidNumber(num string, base int) bool {
	if !isNumber(num) {
		return false
	}
	switch base > 0 {
	case base <= 10:
		for i := range num {
			fmt.Println(num[i])
			fmt.Println(byte(base + 48))
			if int(num[i]) >= base+48 {
				return false
			}
		}

	case base > 10:
		for i := range num {
			if num[i] >= byte(base+65) {
				return false
			}

		}

	}
	return true
}
func main() {
	fmt.Print("Enter P:")
	var p int
	fmt.Scanln(&p)
	r := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(r)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()

	fmt.Println(isValidNumber(s, p))
}
