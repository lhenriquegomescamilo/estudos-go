package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("subtração", a-b)
	fmt.Println("Divisao", a/b)
	fmt.Println("Multiplicacao", a*b)
	fmt.Println("Módulo", a%b)

	// bitwise
	fmt.Println("E => ", a&b)   // 11 & 10 = 10 (binário para decimal = 2)
	fmt.Println("Ou => ", a|b)  // 11 | 10 = 11 (binário para decimal = 3)
	fmt.Println("Xor => ", a^b) // 11 ^ 10 = 01 (binário para decimal = 1)

	c := 3.0
	d := 2.0
	// Outras operações usando math

	fmt.Println("Maior  => ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor  => ", math.Min(c, d))
	fmt.Println("Exponenciação => ", math.Pow(c, d))
}
