package main

import "fmt"

func main() {
	//utiliza esto cuando vas a iterar sobre una misma variable
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
	// Sin condicion
	// cuando quieres anidar multiples condiciones.
	value := 200
	switch {
	case value > 100:
		fmt.Println("el valor es mayor que 100")
	case value < 0:
		fmt.Println("el valor es menor que 0")
	default:
		fmt.Println("we no cumple ninguna condición")
	}
}

//cuando estas ejecutando muntiples funciones if utilizas switch
//en go solo existe el ciclo for
//para comentar
