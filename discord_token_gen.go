package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	now := time.Now()
	start_time_second := now.Second()
	rand.Seed(time.Now().UnixNano())
	fmt.Println("TOKENS\nHow many tokens do you want to generate? - ")
	var amount int
	fmt.Scanf("%d", &amount)
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var tokens = []string{}

	for i := 0; i < amount; i++ {
		// Make the string for "A"
		a := make([]rune, 20)
		for i := range a {
			a[i] = letters[rand.Intn(len(letters))]
		}

		// Make the string for "B"
		b := make([]rune, 6)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}

		// Make the string for "C"
		c := make([]rune, 27)
		for i := range c {
			c[i] = letters[rand.Intn(len(letters))]
		}

		current_Token := "OTIw" + string(a) + "." + string(b) + "." + string(c)
		tokens = append(tokens, current_Token)
	}

	file, err := os.Create("./generated_tokens.txt")
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(file)
	for _, token := range tokens {
		_, err := writer.WriteString(token + "\n")
		if err != nil {
			log.Fatalf("Got error while writing to a file. Err %s", err.Error())
		}
		//fmt.Printf("Bytes Written: %d\n", bytesWritten)
		//fmt.Printf("Available: %d\n", writer.Available())
		//fmt.Printf("Buffered: %d\n", writer.Buffered())

	}

	now2 := time.Now()
	end_time_second := now2.Second()
	fmt.Printf("DONE! -- Wrote %v tokens to file.\nThis took %v seconds", amount, end_time_second-start_time_second)
	writer.Flush()

}
