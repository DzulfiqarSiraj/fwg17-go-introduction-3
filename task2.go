package main

import (
	"fmt"
	"time"
)

func (D DeretBilangan) deretFibonacci() []int {
	results := []int{0, 1}

	f0 := 0
	f1 := 1
	for i := 2; i <= D.N; i++ {
		f2 := f0 + f1
		f0 = f1
		f1 = f2
		if f2 > D.N {
			break
		} else {
			results = append(results, f2)
			time.Sleep(time.Millisecond * 100)
		}
	}

	return results
}

func ganjilGenap(receive chan []int) {
	var deretAngka = []int{}

	deretAngka = <-receive
	for _, value := range deretAngka {
		if value%2 == 0 {
			fmt.Printf("genap %v\n", value)
			time.Sleep(time.Second)
		} else if value%2 != 0 {
			fmt.Printf("ganjil %v\n", value)
			time.Sleep(time.Second)
		}
	}
}
