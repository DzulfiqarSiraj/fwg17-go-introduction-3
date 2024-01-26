package main

import (
	"fmt"
)

func main() {
	DB := DeretBilangan{40}
	D := Deret{40}

	go DB.prima()
	go DB.ganjil()
	go DB.genap()
	go DB.fibonacci()

	fmt.Scanln()

	send := make(chan []int)
	go ganjilGenap(send)
	send <- D.deretFibonacci()

	fmt.Scanln()
}
