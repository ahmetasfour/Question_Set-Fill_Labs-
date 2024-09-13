package main

import "fmt"

func main() {
	/*
	   count := 0

	   	for index := 1; index <= 10; index++ {
	   		count += index
	   	}

	   fmt.Println(count)
	*/
	index := 0

	for index < 10 {

		fmt.Println(index)

		if index == 4 {
			break
		}
		index++
	}

}
