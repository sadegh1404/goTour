package main

import "fmt"

func main() {

	x := 34432

	defer fmt.Println("A", x)

	x = 2

	for i := 0; i < 3; i++ {

		defer fmt.Println("loop", i)

	}
	x = 34432
	defer func() { fmt.Println("B", x) }()

	x = 3

	fmt.Println("end")

}
