package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("Hora atual: %v\n", time.Now().Unix())

	fmt.Println("Aguardando...")

	time.Sleep(2 * time.Second)

	fmt.Printf("Hora atual: %v\n", time.Now().Unix())

}
