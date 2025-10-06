package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   7.98,
		"water": 3.47,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])
}
