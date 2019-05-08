package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"C": {
			"Camilo":    1234,
			"Camilinho": 12345,
		},
		"M": {
			"Mirna":    12345,
			"Mirnoca":  3345,
			"Mirninha": 52345,
		},
	}
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
