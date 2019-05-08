package main

import "fmt"

func main() {
	f1()
	f2("Teste 01", "Teste 02")
	retornoFuncao03, retornoFuncao04 := f3(), f4("teste 01", "teste 02")
	fmt.Println(retornoFuncao03, retornoFuncao04)

	var primeiroRetornoFuncao05, segundoRetornoFuncao05 = f5()

	fmt.Println(primeiroRetornoFuncao05, segundoRetornoFuncao05)

}

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Println("F2", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "F5 - Retorno 1", " F5 - retorno 2"
}
