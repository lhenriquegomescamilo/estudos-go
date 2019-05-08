package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Camilo": 111,
		"Mirna":  222,
		"CAcau":  333,
	}

	funcsESalarios["Camilinho"] = 12336

	fmt.Println(funcsESalarios)

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
