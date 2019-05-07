package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=:", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	data1 := time.Unix(0, 0)

	data2 := time.Unix(0, 0)

	fmt.Println("Data", data1 == data2)
	fmt.Println("Data", data1.Equal(data2))

	type Pessoa struct {
		Nome string
	}

	pessoa1 := Pessoa{"Camilo"}
	pessoa2 := Pessoa{"Camilo"}
	fmt.Println("Pessoas", pessoa1 == pessoa2)

}
