package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	var index int
	fmt.Scanln(&number)
	fmt.Scanln(&index)
	change(number, index)
	fmt.Println()
}

func change(number int, index int) {
	a := strconv.FormatInt(int64(number), 2)
	b := []rune(a)
	if index > len(a) {
		fmt.Println("index is out of range")
		return
	}
	var i int
	if index == 0 {
		i = index
	} else {
		i = index - 1
	}
	if b[i] == 48 {
		b[i]++
	} else {
		b[i]--
	}
	fmt.Println(string(b))
	v, err := strconv.ParseInt(string(b), 2, 64)
	if err != nil {
		return
	}
	fmt.Println(v)
}
