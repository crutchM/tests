package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(22222222222222222)
	b := big.NewInt(22222222222222222)
	fmt.Println(new(big.Int).Add(a, b))
	fmt.Println(new(big.Int).Sub(a, b))
	fmt.Println(new(big.Int).Mul(a, b))
	fmt.Println(new(big.Int).Div(a, b))
}
