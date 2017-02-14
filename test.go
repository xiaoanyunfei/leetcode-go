package main

import (
	"fmt"
)

func CounterFactory(j int) func() int {
	i := j
	return func() int {
		i++
		return i
	}
}

func main() {
	r := CounterFactory(13)
	fmt.Printf("%d\n", r())
	fmt.Printf("%d\n", r())
	fmt.Printf("%d\n", r())
}
