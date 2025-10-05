package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Println("Good Morning", n)
}

func sayBye(n string) {
	fmt.Println("Goodbye", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("jaime")
	sayBye("eli")

	cycleNames([]string{"jaime", "eli", "ivy"}, sayGreeting)

	a1 := circleArea(10.5)
	a2 := circleArea(22.1)
	fmt.Println(a1, a2)
}
