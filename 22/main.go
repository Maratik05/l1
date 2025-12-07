package main

import (
	"fmt"
	"math/big"
)

func main() {
	aStr := "2000000"
	bStr := "3000000"

	a := new(big.Int)
	b := new(big.Int)
	a.SetString(aStr, 10)
	b.SetString(bStr, 10)

	fmt.Printf("a = %s, b = %s\n\n", aStr, bStr)

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)
	quot := new(big.Int).Div(a, b)

	fmt.Printf("Сложение: %s\n", sum)
	fmt.Printf("Вычитание: %s\n", diff)
	fmt.Printf("Умножение: %s\n", prod)
	fmt.Printf("Деление: %s\n", quot)
}
