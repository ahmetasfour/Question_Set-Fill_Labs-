package main

import (
	"fmt"
	"sort"
)

func main() {

	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "csssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	sortedWords := sortWordsBy_A_And_Length(words)

	fmt.Println(sortedWords)
}

func countA(word string) int {
	count := 0
	for _, char := range word {
		if char == 'a' {
			count++
		}
	}
	return count
}

func sortWordsBy_A_And_Length(words []string) []string {

	sort.Slice(words, func(i, j int) bool {

		countA_i := countA(words[i])
		countA_j := countA(words[j])
		if countA_i != countA_j {
			return countA_i > countA_j
		}

		return len(words[i]) > len(words[j])
	})

	return words
}
