package main

import "fmt"

func main() {
	params := []int{2, 4, 6, 8, 10, 11, 12}
	fmt.Println(GetSquares(params))
}

func GetSquares(params []int) int {
	sqCh := make(chan int)
	defer close(sqCh)
	for _, value := range params {
		go func(item int) {
			sqCh <- item * item
		}(value)
	}
	result := 0
	for i := 0; i < len(params); i++ {
		result += <-sqCh
	}

	return result

}
