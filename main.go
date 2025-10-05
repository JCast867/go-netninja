package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	names := []string{"eli", "jaime", "jcast", "ivy"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println(index, value)
	}

	for _, value := range names {
		fmt.Println(value)
	}

}
