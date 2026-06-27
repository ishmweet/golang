package main

import "fmt"

func main() {
	// Signed Integers
	var age int = 18
	var temperature int8 = -10
	var population int64 = 1420000000

	fmt.Println("int:", age)
	fmt.Println("int8:", temperature)
	fmt.Println("int64:", population)

	// Unsigned Integers
	var score uint = 100
	var level uint8 = 255
	var stars uint16 = 65000

	fmt.Println("uint:", score)
	fmt.Println("uint8:", level)
	fmt.Println("uint16:", stars)

	// Floating Point Numbers
	var pi float32 = 3.14
	var gravity float64 = 9.81 // float64 is used in most cases

	fmt.Println("float32:", pi)
	fmt.Println("float64:", gravity)

	// Byte (alias for uint8)
	var letter byte = 'A'
	fmt.Println("byte:", letter)

	// Rune (alias for int32, stores unicode characters)
	var emoji rune = '😊'
	fmt.Println("rune:", emoji)

	// Boolean
	var isgay bool = false
	fmt.Println("bool:", isgay)

	// String
	var name string = "Dexter"
	fmt.Println("string:", name)
}
