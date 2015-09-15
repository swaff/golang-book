package main

import "fmt"
import "golang-book/chapter11/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	min := math.Minimum(xs)
	max := math.Maximum(xs)
	fmt.Println("Average", avg)
	fmt.Println("Minimum", min)
	fmt.Println("Maximum", max)
}
