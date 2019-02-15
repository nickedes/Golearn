package main

import (
	"crypto/sha256"
	"fmt"
)

func Digest() string {
	nonce := 0

	digest := sha256.New()
	s := "text string here!"
	digest.Write([]byte(s))

	for {
		digest.Write([]byte(nonce))
		hashval := fmt.Sprintf("%x\n", digest.Sum(nil))
		if hashval {
			return hashval
		}

	}
}

func main() {
	fmt.Println(Digest())
}
