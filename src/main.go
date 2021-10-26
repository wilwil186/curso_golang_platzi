package main

import "fmt"

func main() { //ciclo tarea repetitiva

	//For condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// For While
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//For forever
	counteForever := 0
	for {
		fmt.Println(counteForever)
		counteForever++
	}
}

//en go solo existe el ciclo for
//para comentar
