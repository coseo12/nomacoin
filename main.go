package main

import (
	"fmt"
	"time"
)

func countToTen(c chan int) {
	for i := range [10]int{} {
		time.Sleep(time.Second * 1)
		fmt.Printf("sending %d\n", i)
		c <- i
	}
	c <- 999
}

func main() {
	// defer db.Close()
	// cli.Start()
	c1 := make(chan int)
	go countToTen(c1)
	for {
		a := <-c1
		fmt.Printf("received %d\n", a)
		if a == 999 {
			break
		}
	}
}
