package main

import (
	"fmt"
	"time"
)

func countToTen(c chan<- int) {
	for i := range [10]int{} {
		time.Sleep(time.Second * 1)
		fmt.Printf("sending %d\n", i)
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for {
		a, ok := <-c
		if !ok {
			fmt.Println("Done!!")
			break
		}
		fmt.Printf("received %d\n", a)
	}
}

func main() {
	// defer db.Close()
	// cli.Start()
	c := make(chan int)
	go countToTen(c)
	receive(c)

}
