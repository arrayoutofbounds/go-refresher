package main

import "fmt"

// var name string = "anmol" // another way to do it

// var name = "ad" // also another way to do it and infer type

func main() {
	name, size := "Anmol", 1.3

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", size)
}
