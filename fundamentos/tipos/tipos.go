package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é ", reflect.TypeOf(32000))

	// Sem sinal (só postiivos)... uint8, unit16, uint32, uint64

	var b byte = 255
	fmt.Println("O valor de byte é ", reflect.TypeOf(b))

	// Com sinal... int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo de um inteiro é", i1)
	fmt.Println("O tipo i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeando da tabela unicode (int32)

	fmt.Println("o rune é ", reflect.TypeOf(i2))

	fmt.Println(i2)

	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal é ", reflect.TypeOf(49.99)) // float64

	// Booleano
	bo := true

	fmt.Println("O tipo de bo é ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "Olá meu nome é Camilo"

	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := `
		Olá 
		meu
		nome 
		é
		Camilo
	`
	fmt.Println("O tamanho da string é ", len(s2))

	//char???
	char := 'a'
	fmt.Println("O tipo de char é ", reflect.TypeOf(char)) // No Go char e reprentado por int32

}
