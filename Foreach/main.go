package main

import "fmt"

func main() {

	//var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 0, 9, 0, 98, 9898}
	/*
		for i := 0; i < len(numbers); i++ {

			fmt.Println(numbers[i])
		}
	*/

	/*
		for index, value := range numbers {

			fmt.Println(index, value)
		}
	*/

	var language = "golang"
	for _, count := range language {
		fmt.Printf("%c\n", count)

	}

}
