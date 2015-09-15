package main

import "fmt"

func fib(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	for i := uint(0); i < 11; i++ {
		fmt.Println("Fib ", i, ":", fib(i))
	}
}
