package main

import (
	"fmt"
	"time"
)

func countToTen(name string) {
	for i := range [10]int{} {
		fmt.Println(i, name)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// defer db.Close()
	// cli.Start()

	go countToTen("first")
	go countToTen("second")
	for {
	}
}
