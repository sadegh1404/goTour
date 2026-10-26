package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Predict befor run
	var u uint8 = 250 // 11111111 MAX IS 255

	u += 10 // 11111111 + 00001010

	fmt.Println(u)

	x := 3.99

	fmt.Println(int(x))

	n := 65

	fmt.Println(string(rune(n)), strconv.Itoa(n))

	var big int64 = 1 << 40 // try with 2 , 16, 32, 31,

	fmt.Println(int32(big))
}
