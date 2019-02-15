package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	s := "text string here!"
	digest := sha256.New()
	digest.Write([]byte("hello world\n"))
	digest.Write([]byte(s))
	fmt.Println(digest.Sum(nil))
	fmt.Printf("%x\n", digest.Sum(nil))
}
