package main

import "fmt"

func collatz(n int) int {
	count := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		count += 1
	}
	return count
}

func collatzRecursive(n int) int {

	if n == 1 {
		return 1
	} else if n%2 == 0 {
		return collatzRecursive(n/2) + 1
	} else {
		return collatzRecursive(3*n+1) + 1
	}
}

func main() {
	maxSteps := 0

	i := 1
	for i < 10000 {
		steps := collatzRecursive(i)
		if steps > maxSteps {
			maxSteps = steps
			fmt.Println(i)
		}
		i++
	}
	fmt.Println(maxSteps)
}
