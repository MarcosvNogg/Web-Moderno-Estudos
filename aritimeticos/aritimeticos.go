package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	b := 3

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtracao =>", b-a)
	fmt.Println("Multiplicacao =>", a*b)
	fmt.Println("Divisao =>", b/a)
	fmt.Println("Módulo =>", b%a)

	// Bitwise - bit a bit
	fmt.Println("And =>", a&b)
	fmt.Println("Or =>", a|b)
	fmt.Println("Xor =>", a^b)

	c := 2.0
	d := 3.0

	// Outras Operacoes usando MATH
	nmax := math.Max(float64(c), float64(d))
	nmin := math.Min(c, d)
	npow := math.Pow(c, d)

	fmt.Println("Max =>", nmax)
	fmt.Println("Min =>", nmin)
	fmt.Println("Pow =>", npow) // Power

}
