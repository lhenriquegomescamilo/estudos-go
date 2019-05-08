package main

import (
	"fmt"
	"time"
)

func main() {
	// Semelhante com o while

	i := 1
	for i <= 10 { // Nào tem while
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\nMisturando estruturas de controle...")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("\nO Número %v é Par", i)
		} else {
			fmt.Printf("\nO número %v é Impar", i)
		}
	}
	for {
		// Láco infinito
		fmt.Println("Para sempre..")
		time.Sleep(time.Second)
	}
}
