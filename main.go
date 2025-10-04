package main

import "fmt"

func main() {

	// strings
	var nameOne string = "jaime"
	var nameTwo = "eli"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "jcast"
	nameThree = "ivy"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Caca"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// ints
	var ageOne int = 22
	var ageTwo = 21
	ageThree := 4

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 256
	fmt.Println(numOne, numTwo, numThree)

	// floats
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 5474547467.67
	scoreThree := 1.5
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
