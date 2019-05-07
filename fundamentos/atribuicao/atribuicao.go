package main

import "fmt"

func main() {
	var b byte = 3

	fmt.Println(b)

	i := 3 // inferencia de tipo

	i += 3 //atribuição aditivos

	i -= 2 // atribuicao substrativa i = i - 2

	i /= 2 // Atribuição de divisão  = i % 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 1, 2

	fmt.Println(x, y)

	x, y = y, x

	fmt.Println(x, y)

}
