package main

import "fmt"

func main() {
	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [5]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Julian"
	fmt.Printf("Students: %v\n", students)
	students[1] = "John"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of students: %v\n", len(students))

	// weird arrays of arrays:

	// U can do that:
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	// or u can do that:
	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}

	fmt.Println(identityMatrix2)

	// make functions
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	b := []int{}
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(b))
	fmt.Printf("Capacity: %v\n", cap(b))

	b = append(b, 1)
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(b))
	fmt.Printf("Capacity: %v\n", cap(b))
}
