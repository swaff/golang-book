package main

import "fmt"

func getGreatest(numbers ...int) int {
	greatest := 0

	for _, i := range numbers {
		if i > greatest {
			greatest = i
		}
	}
	return greatest
}

func main() {
	numbers := []int{1, 40, 299, 9994, 29, 299}
	fmt.Println("The greatest number in [1, 40, 299, 9994, 29, 299] is:", getGreatest(numbers...))
}
