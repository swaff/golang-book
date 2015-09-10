package main

import "fmt"

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celcius := (fahrenheit - 32) * 5 / 9
	fmt.Println(fahrenheit, "is", celcius, "celcius")
}
