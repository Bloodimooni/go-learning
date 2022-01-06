package main

import (
	"fmt"
	"math"
)

func main() {
	// main function

	fmt.Println("\n", square(15125), "\n", root(64))

	for i := 0; i < 200000; i++ {
		fmt.Println(byte_to_string(rune(i)), " ", i)
	}

}

func square(value float64) float64 {
	return value * value
}

func root(value float64) float64 {
	return math.Sqrt(value)
}

func byte_to_string(value rune) string {
	return string(value)
}
