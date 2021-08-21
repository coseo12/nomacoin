package main

import (
	"fmt"
	"time"
)

func send(c chan<- int) {
	for i := range [10]int{} {
		fmt.Printf("sending %d\n", i)
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for {
		time.Sleep(time.Second * 1)
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
	c := make(chan int, 10)
	go send(c)
	receive(c)

}
