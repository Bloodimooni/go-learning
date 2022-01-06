package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("names.txt")
	if err != nil {
		fmt.Println("Error reading", err)
		return
	}

	var current_name string
	var names string

	for i := 0; i < len(data); i++ {
		// fmt.Println(data[i])
		if data[i] != 44 {
			current_name += string(data[i])
		}

		if data[i] == 44 {
			names += " " + current_name
			current_name = ""
		}
	}

	fmt.Println(names)
	// all names are laoded from file, still takes a bit too long
}
