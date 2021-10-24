package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.14
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	//Declaración de variables enteras
	//coloca los dos puntos cuando la variable no a sido creada con anterioridad
	base := 12
	//otra forma es
	var altura int = 14
	//otra forma
	var area int

	fmt.Println(base, altura, area) //se necesita utulizar toda variable que se declara.
	//Zero values.
	//por defecto estos valores tienen.
	var a int     //0
	var b float64 //0
	var c string  //string vacio
	var d bool    // false
	fmt.Println(a, b, c, d)
	//ejercicio Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)

}

//para comentar
