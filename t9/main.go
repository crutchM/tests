package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	go genData(arr, ch1)
	go powData(ch1, ch2)
	for x := range ch2 {
		fmt.Println(x)
	}

}

func genData(source []int, output chan int) {
	for i := 0; i < len(source); i++ {
		output <- source[i]
	}
	close(output)
}

func powData(input chan int, output chan int) {
	for x := range input {
		output <- x * 2
	}
	close(output)
}
