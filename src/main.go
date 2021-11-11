package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}

func main() {
	// este paquete te permite interartuar con las gorutines
	var wg sync.WaitGroup

	fmt.Println("Hello")
	//vamos a agregar una gorutine
	wg.Add(1)
	//para caplicar las gorutines
	go say("world", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")
	// espera we
	time.Sleep(time.Second * 1)

}
