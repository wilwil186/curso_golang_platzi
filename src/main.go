package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2

}

func main() {

	a := 50
	b := &a

	fmt.Println(b)
	// el * es la operación contraria de &
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	//instanciar
	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)
	//retornar un mensaje aletorio
	myPC.ping()
	//estado actual
	fmt.Println(myPC)
	//suplicar
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)

}
