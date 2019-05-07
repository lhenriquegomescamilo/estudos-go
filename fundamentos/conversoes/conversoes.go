package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9

	notaFinal := int(nota)

	fmt.Println(notaFinal)

	//cuidado...
	fmt.Println("Representa o código da tabela ASCII " + string(97))

	// int para string

	fmt.Println("Teste conversao de inteiro para string " + strconv.Itoa(123))

	// String to Int

	num, _ := strconv.Atoi("123")

	//
	fmt.Println("Conversao para string para inteiro", (num - 122))

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro")
	}

}
