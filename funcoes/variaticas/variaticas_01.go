package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}
func main() {
	fmt.Printf("Média; %.2f", media(1, 2, 3, 4, 5, 9, 10, 20, 1.192, 3.343, 100.12))
}
