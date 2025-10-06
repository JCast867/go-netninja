package main

import (
	"fmt"
)

func updateName(x string) {
	x = "wedge"
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	name := "jaime"
	updateName(name)
	fmt.Println(&name)
	m := &name
	fmt.Println(m)
	fmt.Println(*m)
	fmt.Println(name)
}
