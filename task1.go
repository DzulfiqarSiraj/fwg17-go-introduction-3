package main

import (
	"fmt"
	"time"
)

type DeretBilangan struct {
	N int
}

func (D DeretBilangan) prima() {
	for i := 2; i <= D.N; i++ {
		bilPrima := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				bilPrima = false
				break
			}
		}
		if bilPrima {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (D DeretBilangan) ganjil() {
	for i := 1; i <= D.N; i++ {
		if i%2 != 0 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (D DeretBilangan) genap() {
	for i := 1; i <= D.N; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (D DeretBilangan) fibonacci() {
	f0 := 0
	f1 := 1
	fmt.Println(f0)
	fmt.Println(f1)
	for i := 2; i <= D.N; i++ {
		f2 := f0 + f1
		f0 = f1
		f1 = f2
		if f2 > D.N {
			break
		} else {
			fmt.Println(f2)
			time.Sleep(time.Millisecond * 100)
		}
	}
}
