package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := x[0]
	largest := x[0]

	for _, value := range x {
		if value < smallest {
			smallest = value
		}

		if value > largest {
			largest = value
		}
	}

	fmt.Println("The smallest number in the list is:", smallest)
	fmt.Println("The largest number in the list is:", largest)
}
