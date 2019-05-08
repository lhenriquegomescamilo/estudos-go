package main

import "fmt"

func main() {
	// Mapas devem ser inicializados
	//var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[123456789] = "Mirna"
	aprovados[1231232132132] = "Camilo"
	aprovados[123123213] = "Maicon"
	aprovados[123122132] = "Suzy"
	aprovados[15632132132] = "Otavio"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s  (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[1231232132132])

	delete(aprovados, 1231232132132)
	fmt.Println(aprovados)

}
