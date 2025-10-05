package main

import "fmt"

func main() {
	age := 22
	name := "jaime"

	// Print
	fmt.Print("Hello\n")

	// Println
	fmt.Println("World!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted sring) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v\n", age, name)
	fmt.Printf("my age is %q and my name is %q\n", age, name) // %q puts quotes around the string
	fmt.Printf("age is of type %T\n", age)                    // %T gives you the type of the variable
	fmt.Printf("you scored %f points!\n", 225.55)             // %f is for floats and you're allowed to put 0.1, 0.01, etc to round the decimal points
	fmt.Printf("you scored %0.1f points!\n", 225.55)

	// Sprintf (save formatted atrings)
	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println("Saved string is:", str)
}
