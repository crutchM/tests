package main

import (
	"fmt"
	"strings"
)

func reverse(row string) string {
	splited := strings.Split(row, " ")
	reversed := make([]string, len(splited))
	for i := 0; i < len(splited); i++ {
		reversed[i] = splited[len(splited)-1-i]
	}
	return strings.Join(reversed[:], " ")
}

func main() {
	words := "snow dog sun"
	wordsRu := "снег собака солнце"
	fmt.Println(reverse(words))
	fmt.Println(reverse(wordsRu))

}
