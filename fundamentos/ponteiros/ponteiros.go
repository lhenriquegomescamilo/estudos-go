package main

import "fmt"

func main() {
	i := 1

	// Go não possui aritméticas de ponteiros
	var p *int = nil

	p = &i          // Recuperando o endereço da variável i
	fmt.Println(p)  // Apresentando o endereço da variável
	fmt.Println(*p) // Apresentando o conteúdo do endereço

	// Alterando o conteúdo do ponteiro
	*p++
	fmt.Println(*p)

	// Alterando a própria variável
	i++
	fmt.Println(i)

	fmt.Println(p, *p, i, &i) // Exibindo o conteúdo da variável e o conteúdo do ponteiro

}
