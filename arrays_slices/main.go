package main

import "fmt"

func main() {
	var names [3]string
	fmt.Println("Array length:", len(names)) // Print the array length
	fmt.Println("Before assignment:", names)

	names[0] = "asf"
	names[1] = "mrt"
	names[2] = "kos"

	fmt.Println("After assignment:", names)
	fmt.Println("Third element:", names[2])
}
