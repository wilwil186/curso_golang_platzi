package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("es verdad")
	}
	// With or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("es verdad, oki doki")
	}
}

//en go solo existe el ciclo for
//para comentar
