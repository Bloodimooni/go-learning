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

	//	var name_to_compare string
	var current_name []byte
	var name_list []byte
	var names []byte
	names = data
	// load all the names into the list "name_list"

	for i := 0; i < len(names); i++ {

		if string(names[i]) != "," {
			current_name = append(current_name, names[i])
		}

		if string(names[i]) == "," {
			current_name = append(current_name, 13)
			current_name = append(current_name, 10)
			for e := 0; e < len(current_name); e++ {
				name_list = append(name_list, current_name[e])
			}
			current_name = nil
		}
	}

	// sort out the list before I write the file:
	//os.WriteFile("sorted.txt", name_list, 0644)
	// I want to take a word and save it, then make sure its nowhere else in the file.
	// so what I can do is I could take the word and just compare it to everything in the file.
	// this should print out all name without the commas
}

// thought the program didn't work cause it finished in 0.7 seconds. -(°<0>°)-
