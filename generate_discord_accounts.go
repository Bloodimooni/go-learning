package main

import (
	"fmt"
	"io/ioutil"
)

const (
	registerSite = "https://discord.com/register"
	emailSite    = "https://10minutemail.net/?lang=en"
)

func main() {

}

func get_names() {
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

}
