package main

import "fmt"

func main() {
	var base int
	var num string
	fmt.Println("Enter numbers")
	fmt.Scan(&num)
	fmt.Println("Enter base")
	fmt.Scan(&base)
	fmt.Println(isValidNumber())
	return
}

func isValidNumber(num string, base int) bool {
	for i := range num {
		if num[i] >= base {
			return false
		}else if num[i] =! 1 || 2 || 0 || 3 || 4|| 5 || 6 || 7 || 8 || 9{
			if num[i] >= base {	
				return false
			}
		}else {
			return true
		}
	}
	return true
}
