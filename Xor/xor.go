package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sample(A []int, x int) []int {
	rand.Seed(time.Now().UnixNano())
	randomize := rand.Perm(len(A))

	sampleslice := make([]int, 0)

	for _, v := range randomize[:x] {
		sampleslice = append(sampleslice, A[v])
	}
	return sampleslice
}

func main() {
	fmt.Println(1)
}
