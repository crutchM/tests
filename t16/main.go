package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	fmt.Println(ar)
	fmt.Println("after sort")
	fmt.Println(quickSort(ar))

}

func quickSort(base []int) []int {
	if len(base) < 2 {
		return base
	}
	leftEdge, rightEdge := 0, len(base)-1
	//берем рандомную точку опоры
	pivot := rand.Int() % len(base)
	base[pivot], base[rightEdge] = base[rightEdge], base[pivot]

	for i := range base {
		if base[i] < base[rightEdge] {
			base[i], base[rightEdge] = base[rightEdge], base[i]
		}
	}
	base[leftEdge], base[rightEdge] = base[rightEdge], base[leftEdge]
	quickSort(base[:leftEdge])
	quickSort(base[leftEdge+1:])
	return base
}
