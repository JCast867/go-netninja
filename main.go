package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greeting := "hello there !"

	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	ages := []int{45, 20, 78, 53, 25, 12, 56}
	sort.Ints(ages)
	fmt.Println((ages))

	index := sort.SearchInts(ages, 53)
	fmt.Println(index)

	names := []string{"eli", "jaime", "jcast", "ivy"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "eli"))
}
