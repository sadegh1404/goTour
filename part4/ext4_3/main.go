package main

import (
	"errors"
	"fmt"
)

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func Describe(v any) string {
	switch x := v.(type) {
	case int:
		parity := "odd"
		if x%2 == 0 {
			parity = "even"
		}
		return fmt.Sprintf("int: %d (%s)", v, parity)
	case string:
		return fmt.Sprintf("string of %d runes ", len([]rune(x)))
	case []int:
		return fmt.Sprint("slice with sum  ", sum(x))
	case error:
		return fmt.Sprintf("error: %s", v)
	case nil:
		return "nil"
	case fmt.Stringer:
		return "Its String()"
	default:
		return fmt.Sprintf("unknown %T", x)
	}
}

func main() {
	fmt.Println(Describe(42))
	fmt.Println(Describe(7))
	fmt.Println(Describe("hello"))
	fmt.Println(Describe("سلام"))
	fmt.Println(Describe([]int{1, 2, 3}))
	fmt.Println(Describe(errors.New("disk full")))
	fmt.Println(Describe(nil))
	fmt.Println(Describe(3.14))
	fmt.Println(Describe(true))

}
