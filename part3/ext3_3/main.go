package main

import "fmt"

func removeAt(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func removeFast(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func main() {
	s := []string{"a", "b", "c", "d"}
	r := removeAt(s, 1)
	fmt.Println(r, s) // predict before running!

	s2 := []string{"a", "b", "c", "d"}
	f := removeFast(s2, 1)
	fmt.Println(f, s2)
}
