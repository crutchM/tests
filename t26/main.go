package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(checkString(input))
}

func checkString(input string) bool {
	dict := make(map[string]int)
	for _, v := range input {
		dict[string(v)]++
	}
	for _, v := range dict {
		if v > 1 {
			return false
		}
	}
	return true
}
