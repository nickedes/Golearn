package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to channel c
}

func main() {
	go say("world")
	say("hello")

	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // create channel c
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from channel c

	fmt.Println(x, y, x+y)

}
