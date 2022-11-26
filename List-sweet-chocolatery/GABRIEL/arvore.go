package main

import "fmt"

type Creature struct {
	Species string
}

func (c Creature) Reset() {
	c.Species = ""
}

func main() {

	var creature string = "shark"
	var pointer *string = &creature

	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)

	fmt.Println("*pointer =", *pointer)

	*pointer = "jellyfish"
	fmt.Println("*pointer =", *pointer)

	fmt.Println("creature =", creature)

}
