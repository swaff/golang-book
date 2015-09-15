package main

import "fmt"

func swap(xPtr *int, yPtr *int) {
	temp := *xPtr
	*xPtr = *yPtr
	*yPtr = temp
}

func main() {
	x := 1
	y := 2

	fmt.Println("X is ", x)
	fmt.Println("Y is ", y)
	fmt.Println("swapping...")
	swap(&x, &y)
	fmt.Println("X is ", x)
	fmt.Println("Y is ", y)
}
