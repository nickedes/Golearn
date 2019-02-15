package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

func Digest() string {
	nonce := 0

	d := 4
	x := ""
	for i := 0; i < d; i++ {
		x += "0"
	}

	digest := sha256.New()
	s := "t1ext string here!"
	digest.Write([]byte(s))

	for {
		digest.Write([]byte(strconv.Itoa(nonce)))
		hashval := fmt.Sprintf("%x\n", digest.Sum(nil))
		if strings.HasPrefix(hashval, x) {
			return hashval
		}
		nonce += 1
	}
}

func main() {
	fmt.Println(Digest())
}
