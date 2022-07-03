package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(remove(arr, 3))
}

func remove(arr []int, index int) []int {
	index--
	return append(arr[:index], arr[index+1:]...)
}
