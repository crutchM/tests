package main

import "fmt"

func reverse(row string) string {
	rowRune := []rune(row)
	reversedRune := make([]rune, len(rowRune))
	for i := 0; i < len(rowRune); i++ {
		reversedRune[i] = rowRune[len(rowRune)-1-i]
	}
	return string(reversedRune)
}

func main() {
	firstRow := "12efasdmkm122k0392i540934utoiewjfomasd"
	secondRow := "тевирп"
	fmt.Println(reverse(firstRow))
	fmt.Println(reverse(secondRow))
}
