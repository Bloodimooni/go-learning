package primitives

import (
	"fmt"
)

func main() {
	fmt.Println("primitives:")

	// boolean data:

	var n bool = false
	fmt.Printf("%v, %T", n, n)

	a := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", m, m)

	var b bool
	fmt.Printf("%v, %T\n", b, b)

	//// neumaric types:

	// signed integers:
	c := 42
	fmt.Printf("%v, %T \n", c, c)

	//unsigned integers: 	They go from 0 to byte size !!!! NO negative VALUES !!!!
	var d uint16 = 42
	fmt.Printf("%v, %T \n", d, d)

	// Bit operators:

	e := 10
	f := 3

	fmt.Println(e & f)  // and
	fmt.Println(e | f)  // or
	fmt.Println(e ^ f)  // xor
	fmt.Println(e &^ f) // xand

	// bit shifting:
	g := 8              // 00000100
	fmt.Println(g << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(g >> 3) // 2^3 / 2^3 = 2^0

	// floating types:
	var j float32 = 3.14
	j = 13.7e12
	fmt.Printf("%v, %T\n", j, j)

	// COMPLEX NUMBERS:
	var k complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", real(k), real(k))
	fmt.Printf("%v, %T\n", imag(k), imag(k))

}
