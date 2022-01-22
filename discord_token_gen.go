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

	fmt.Printf("Okay, I will generate %v tokens.", amount)
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
	fmt.Printf("DONE! -- Wrote %v tokens to file.\nThis took %v seconds\n\n\n", amount, end_time_second-start_time_second)
	writer.Flush()

	//for y := 0; y < len(tokens); y++ {

	//cfuid := make([]rune, 27)
	//for i := range cfuid {
	//	cfuid[i] = letters[rand.Intn(len(letters))]
	//}

	//dcfduid := make([]rune, 27)
	//for i := range dcfduid {
	//	dcfduid[i] = letters[rand.Intn(len(letters))]
	//}

	//var header = map[string]string{
	//	"authority":            "discord.com",
	//	"scheme":               "https",
	//	"accept":               "*/*",
	//	"accept-encoding":      "gzip, deflate, br",
	//	"accept-language":      "en-US",
	//	"Authorization":        "value",
	//	"content-length":       "0",
	//	"cookie":               "cookie",
	//	"referer":              "https://discord.com/channels/@me",
	//	"sec-fetch-dest":       "empty",
	//	"sec-fetch-mode":       "cors",
	//	"sec-fetch-site":       "same-origin",
	//	"user-agent":           "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/37.0.2062.94 Chrome/37.0.2062.94 Safari/537.36",
	//	"x-context-properties": "eyJsb2NhdGlvbiI6Ikludml0ZSBCdXR0b24gRW1iZWQiLCJsb2NhdGlvbl9ndWlsZF9pZCI6Ijg3OTc4MjM4MDAxMTk0NjAyNCIsImxvY2F0aW9uX2NoYW5uZWxfaWQiOiI4ODExMDg4MDc5NjE0MTk3OTYiLCJsb2NhdGlvbl9jaGFubmVsX3R5cGUiOjAsImxvY2F0aW9uX21lc3NhZ2VfaWQiOiI4ODExOTkzOTI5MTExNTkzNTcifQ==",
	//	"x-super-properties":   "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiRGlzY29yZCBDbGllbnQiLCJyZWxlYXNlX2NoYW5uZWwiOiJjYW5hcnkiLCJjbGllbnRfdmVyc2lvbiI6IjEuMC42MDAiLCJvc192ZXJzaW9uIjoiMTAuMC4yMjAwMCIsIm9zX2FyY2giOiJ4NjQiLCJzeXN0ZW1fbG9jYWxlIjoic2siLCJjbGllbnRfYnVpbGRfbnVtYmVyIjo5NTM1MywiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=",
	//}
	//header["cookie"] = "__cfuid=" + string(cfuid) + "; __dcfduid=" + string(dcfduid) + "; locale=en-US"
	//header["Authorization"] = tokens[y]
	//}

}
