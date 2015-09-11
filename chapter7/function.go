package main

import "fmt"

func average(numbers []float64) float64 {
	var total float64 = 0

	for _, value := range numbers {
		total += value
	}
	return total / float64(len(numbers))
}

func main() {
	numbers := []float64{10, 33, 73, 73, 24}
	fmt.Println("Average:", average(numbers))
}
