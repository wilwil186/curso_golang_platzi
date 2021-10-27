package main

import "fmt"

func main() {
	//Defer
	defer fmt.Println("hola")
	fmt.Println("mundo")
	//continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		// Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		// Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}

//cuando estas ejecutando muntiples funciones if utilizas switch
//en go solo existe el ciclo for
//para comentar
