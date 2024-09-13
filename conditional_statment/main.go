package main

import "fmt"

func main() {
	/*
		var age = 17
		if age >= 18 {
			fmt.Println("you can vote")
		} else {
			fmt.Println("you cant vote")
		}
	*/

	a := 23146
	b := 6465345
	c := 3345
	if a >= b && a >= c {

		fmt.Println("a is Tha Big: ", a)

	} else if b >= a && b >= c {
		fmt.Println("b is Tha Big: ", b)
	} else {
		fmt.Println("c is Tha Big: ", c)
	}
}
