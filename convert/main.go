package main

import (
	"fmt"
)

func Convert(str string, base int) {
	num := 0
	pow := 1
	for i := 0; i == len(str); i++ {
		num = string(str[i]) * pow
		pow = pow * base
	}
	return num
}

func main() {
	fmt.Println("convorted;", Convert(102, 3))
}
