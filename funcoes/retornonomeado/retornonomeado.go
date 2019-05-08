package main

import "fmt"

func main() {
	var x, y = trocar(10, 20)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1

	return // retorno limpo
}
