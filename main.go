package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// removing the pointer receiver (*) results in the main function printing the original value of the Vertex Abs()
// from A Tour of Go: Methods with pointer receivers can modify the value to which the receiver points. Since
// methods often need to modify their receiver, pointer receivers are more common than value receivers.

func (v *Vertex) Scale(f float64) {
	fmt.Println("Vertex before scale: ", v)
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("Vertex after scale: ", v)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println("Print the vertex in main: ", v)
	fmt.Println("Print the abs of vertex in main: ", v.Abs())
}
