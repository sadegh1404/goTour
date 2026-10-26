package main

import "fmt"

func CToF(c float64) float64 {
	f := (c * 9 / 5) + 32
	return f
}

func FToC(f float64) float64 {
	c := (f - 32) / 1.8
	return c
}

func main() {
	fmt.Print("Celsius   |   Fahrenheit\n")
	for c := -10; c <= 40; c += 10 {
		fmt.Printf("%.1f°C   |   %.1f°F\n", float64(c), CToF(float64(c)))
	}

}
