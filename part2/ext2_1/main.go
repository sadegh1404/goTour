package main

import "fmt"

func classic() {
	for i := 1; i <= 30; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func whileLoop() {
	i := 1
	for i <= 30 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
		i++
	}
}

func switchLoop() {

	i := 1
	for i <= 30 {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
		i++
	}

}

func main() {
	fmt.Println("_________ Classic _________")
	classic()
	fmt.Println("_________ WhileLoop _________")
	whileLoop()
	fmt.Println("_________ Switch _________")
	switchLoop()
}
