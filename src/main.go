package main

import "fmt"

func main() {
	//en muchos lenguajes de progamación existen las estructuras de datos llave valor
	//para acceder a un valor necesitas una llave
	//en python son diccionario y en go maps
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
