package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	fmt.Println("Start!")
	rsakey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(rsakey)
	} else {
		fmt.Println("Rsa key generation Error!")
	}
}
