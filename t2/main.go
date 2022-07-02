package main

import (
	"fmt"
	"sync"
)

func main() {
	params := []int{2, 4, 6, 8, 10, 11, 12}
	GetSquares(params)
}

func GetSquares(params []int) {
	var wg sync.WaitGroup
	wg.Add(len(params))
	for _, value := range params {
		go func(item int) {
			defer wg.Done()
			fmt.Println(item * item)
		}(value)
	}
	wg.Wait()
}
