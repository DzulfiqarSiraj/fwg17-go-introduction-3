package main

import (
	"fmt"
)

func main() {

	D := DeretBilangan{40}

	go D.prima()
	go D.ganjil()
	go D.genap()
	go D.fibonacci()

	fmt.Scanln()

	send := make(chan []int)
	go ganjilGenap(send)
	send <- D.deretFibonacci()

	fmt.Scanln()
}
