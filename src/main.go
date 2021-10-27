package main

import "fmt"

func main() {
	//Array
	var Array [4]int
	Array[0] = 1
	Array[1] = 2
	fmt.Println(Array, len(Array), cap(Array))

	//slide
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap((slice)))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])
	//no se puede añadir elementos sen los array
	//por que los arrays son inmutables y los dlide no
	//para eso utilizamos el comando Append
	slice = append(slice, 7)
	fmt.Println(slice)
	//cuando necesitamos agregar una lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
	//arrays inmutables, slice no inmutables sin embargo ya que los arrays son inmutables son mas eficientes
}

//cuando estas ejecutando muntiples funciones if utilizas switch
//en go solo existe el ciclo for
//para comentar
