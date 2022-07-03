package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(set(sequence))
}

func set(input []string) (result []string) {
	newSet := make(map[string]bool)
	for _, v := range input {
		newSet[v] = true
	}
	for val := range newSet {
		result = append(result, val)
	}
	return
}
