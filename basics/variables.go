package main

import "fmt"

func main() {

	// Variable Declaration
	var name string = "Dexter"
	fmt.Println(name)

	// Type Inference
	var age = 35
	fmt.Println(age)

	// Short Variable Declaration
	city := "Miami"
	fmt.Println(city)

	// Multiple Variables
	var a, b, c = 1, 2, 3
	fmt.Println(a, b, c)

	// Zero Values
	var number int
	var isStudent bool
	var language string

	fmt.Println(number)
	fmt.Println(isStudent)
	fmt.Println(language)

	// Updating Variables
	kills := 90
	fmt.Println(kills)

	kills = 100
	fmt.Println(kills)
}
