package main

import "fmt"

func main() {
	fmt.Print("Entre um número: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}