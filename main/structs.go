package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

// Definindo métodos
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}

	// Acessando fields
	fmt.Printf("%f\n", c.r)

	// Acessando métodos
	fmt.Printf("%f\n", c.area())
}
