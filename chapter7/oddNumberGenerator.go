package main

import "fmt"

// Creates a function that returns a function
// The returned function will use a closure to maintain
// a known value, each time the returned generator function
// is called it increments the counter to create the next
// odd number
func makeOddNumberGenerator() func() uint {
	i := uint(1)
	return func() uint {
		odd := i
		i += 2
		return odd
	}
}

func main() {
	generator := makeOddNumberGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
