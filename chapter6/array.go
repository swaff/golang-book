package main

import "fmt"

func main() {
	numbers := [5]float64{10, 33, 73, 73, 24}
	var total float64 = 0
	var average float64 = 0

	for _, value := range numbers {
		total += value
	}
	average = total / float64(len(numbers))

	fmt.Println("Total:", total, "Average:", average)
}
