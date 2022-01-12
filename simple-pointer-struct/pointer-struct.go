// Golang program to show how to instantiate
// Struct Pointer Address Operator
package main

import "fmt"

type shape struct {
	length  int
	breadth int
	color   string
}

func main() {

	// Passing all the parameters
	var shape1 = &shape{10, 20, "Green"}

	// Printing the value struct is holding
	fmt.Println(shape1)

	// Passing only length and color value
	var shape2 = &shape{}
	shape2.length = 10
	shape2.color = "Red"

	// Printing the value struct is holding
	fmt.Println(shape2)

	// Printing the length stored in the struct
	fmt.Println(shape2.length)

	// Printing the color stored in the struct
	fmt.Println(shape2.color)

	// Passing only address of the struct
	var shape3 = &shape{}

	// Passing the value through a pointer
	// in breadth field of the variable
	// holding the struct address
	(*shape3).breadth = 10

	// Passing the value through a pointer
	// in color field of the variable
	// holding the struct address
	(*shape3).color = "Blue"

	// Printing the values stored
	// in the struct using pointer
	fmt.Println(shape3)
	
}
