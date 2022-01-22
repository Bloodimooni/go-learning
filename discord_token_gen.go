package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("TOKEN GENERATOR\nHow many tokens do you want to generate? - ")
	reader := bufio.NewReader(os.Stdin)
	option, _ := reader.ReadString('\n')
	amount, _ := strconv.ParseInt(option, 10, 64)
	fmt.Printf("%T", amount)

}
