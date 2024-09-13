package main

import "fmt"

func findDivisibleNumbers(n int) {
	if n > 2 {
		findDivisibleNumbers(n / 2)
	}

	fmt.Println(n)
}

func main() {
	findDivisibleNumbers(9)
}
