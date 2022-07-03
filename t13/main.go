package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println("a=", a, "b=", b)
	a, b = b, a
	fmt.Println("a=", a, "b=", b)
}
