package main

import "fmt"

func main() { //ejercicio Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)

	x := 10
	y := 58

	//suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//multiplicación
	result = x * y
	//División
	result = y / x
	fmt.Println("División:", result)
	//modulo
	result = y % x
	fmt.Println("Modulo:", result)
	//Incremental
	x++
	fmt.Println("Incremental:", x)
	//Decremental}
	x--
	fmt.Println("Decremental:", x)

}

//para comentar
