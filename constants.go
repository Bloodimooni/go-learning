package main

import "fmt"

const (
	a = iota
	b // since we have iota, it thinks we have a pattern and its also IOTA
	c
)

const (
	a2 = iota // iota is set to a constant block!
)

const (
	_             = iota + 5 // to compare if const is 0 do errorconstant = iota
	catSpecialist = 1 << (10 * iota)
	dogSpecialist
	snakeSpecialist
)

func main() {
	// constants
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", a2, a2)

	fmt.Printf("%v\n", catSpecialist) // --> returns true if we have: var specialistType int = dogSpecialist // IT STILL COMPARES THE TYPE NOT THE VALUE!
	filesize(4000000000.)
	perms()
}

// example for bitshifting:

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func filesize(size float64) {
	fileSize := size
	fmt.Printf("%.2fGB \n", fileSize/GB)
}

// another example : permissions --->

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func perms() {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	// we encoded 3 user roles into a single byte of data!!! thats valuable!
	fmt.Printf("Is Admin? - %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is Headquarters? - %v\n", isHeadquarters&roles == isHeadquarters)
}
