package main

import (
	"fmt"
)

func mostFrequent(data []string) string {
	frequency := make(map[string]int)

	for _, item := range data {
		frequency[item]++
	}

	var mostCommon string
	maxCount := 0

	for key, value := range frequency {
		if value > maxCount {
			maxCount = value
			mostCommon = key
		}
	}

	return mostCommon
}

func main() {
	data := []string{"apple", "pie", "apple", "red", "red", "red"}

	result := mostFrequent(data)

	fmt.Println(result)
}
