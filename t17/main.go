package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	arr := []int{4, 7, 8, 9, 15, 28, 33, 44, 68}
	res, err := binarySearch(arr, 9)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
	fmt.Println(arr[res])
}

func binarySearch(array []int, target int) (int, error) {
	low, high := 0, len(array)-1
	for low <= high {
		mid := (low + high) / 2
		res := array[mid]
		if res == target {
			return mid, nil
		}
		if res > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, errors.New("can't find target")

}
