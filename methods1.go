//A simple program to calculate area and circumference using Methods in Go

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func (c Circle) calculateArea() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) calculateCircumf() float64 {
	return 2 * math.Pi * c.r
}

func main() {

	area := Circle{
		r: 2,
	}
	fmt.Println("Area: ", area.calculateArea())
	fmt.Println("Circumference: ", area.calculateCircumf())
}
