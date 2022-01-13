package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Type in the string I shall copy:")
	reader := bufio.NewReader(os.Stdin)
	option, _ := reader.ReadString('\n')
	fmt.Println("U wrote:\n", option)
}
