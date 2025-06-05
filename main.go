package main

import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	wordMap := make(map[string]int)
	for _, message := range messages {
		wordSet := strings.SplitSeq(message, " ")
		for word := range wordSet {
			if word == "" {
				continue
			}
			wordMap[strings.ToLower(word)]++
		}
	}
	fmt.Println("\n", wordMap)
	return len(wordMap)
}
