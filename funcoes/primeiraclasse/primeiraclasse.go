package main

import "fmt"

func main() {
	var soma = func(a, b int) int {
		return a + b
	}

	fmt.Println(soma(2, 3))

	subtracao := func(a, b int) int {
		return a - b
	}
	fmt.Println(subtracao(2, 1))
}
