package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func double(nums []int) {
	for i := 0; i < len(nums); i++ {
		nums[i] *= 2
	}
}

func addOne(nums []int) {
	nums = append(nums, 99)
}

func main() {
	a := 2
	b := 3
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println("Break")

	nums := []int{2, 3, 4, 5, 6}
	double(nums)
	fmt.Println(nums)
	fmt.Println(&nums)

	s := []int{1, 2, 3}
	addOne(s)
	fmt.Println(s)
}
