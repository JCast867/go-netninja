package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 25, 22}
	var ages = [3]int{20, 25, 22}

	names := [4]string{"eli", "jaime", "jcast", "ivy"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	rangeOne = append(rangeOne, "bob")
	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
