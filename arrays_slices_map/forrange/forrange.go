package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // Quem ira contar é o compilador
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}

	for indices := range numeros {
		fmt.Println(&numeros[indices])
	}
}
