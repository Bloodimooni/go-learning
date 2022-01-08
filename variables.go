package main

import (
	"fmt"
	"strconv"
)

var (
	i            int     = 42
	actorName    string  = "Actor"
	companion    string  = "amongus"
	doctorNumber int     = 15125
	time_spent   float64 = 0.1516

	/// uppercase variables are PUBLIC variable
	/// lowercase variables are only visible in the same package
)

func main() {
	// variables in a function are only visible in the same function (except we return it)
	var sus int = 2516
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", actorName, actorName)
	fmt.Printf("%v, %T\n", companion, companion)
	fmt.Printf("%v, %T\n", doctorNumber, doctorNumber)
	fmt.Printf("%v, %T\n", time_spent, time_spent)
	fmt.Printf("%v, %T\n", sus, sus)

	// heres how to really convert to a string

	var number int = 42
	fmt.Println(string(number), " --- string conversion gives out the Ascii character for 42")
	fmt.Println(strconv.Itoa(number))
}
