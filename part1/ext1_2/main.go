package main

import "fmt"

func divmod(a, b int) (q, r int) {
	q = a / b
	r = a - (q * b)
	// it returns the named value in retrun type of function
	return
}

func swap3(a, b, c int) (int, int, int) {
	return c, b, a
}

func main() {
	fmt.Println(divmod(17, 5))
	fmt.Println(swap3(1, 2, 3))
}
