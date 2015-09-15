package main

import "fmt"

func halfAndIsEven(number int) (int, bool) {
	half := number / 2
	isEven := number%2 == 0
	return half, isEven
}
func main() {

	halfOne, isOneEven := halfAndIsEven(1)
	fmt.Println("1", halfOne, isOneEven)
	halfFour, isFourEven := halfAndIsEven(4)
	fmt.Println("4", halfFour, isFourEven)
}
