package main

import "fmt"

func main() {
	fmt.Println(obterResultado(6.2))
}

// Não tem operadores ternários
func obterResultado(nota float64) string {
	//return nota >=6 ? "Aprovado" : "Reprovado" Infelizmente não existe
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

}
