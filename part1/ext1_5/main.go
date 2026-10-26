package main

import "fmt"

func main() {
	const Big = 1 << 10

	const Small = Big >> 99

	fmt.Println(Big)
	fmt.Println(Small)
}
