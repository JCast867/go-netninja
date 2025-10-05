# Go Programming Language Tutorial
## 1. Intro & Setup

### What is Go?
* Go is a fast, statically types, compiled language. Statically types means that variables types must be declared
* Go is a genereal purpose language
* Go has a built-in testing suppoer
* Go IS an object-oriented language (in its own way)

## 2. Your First Go File
Not much to see here. Just uploaded `test.go` to see it ig

## 3. Variables, Strings, and Numbers
### Variables
There are some key things to note about variables and defining them before we get into variable types
1. If a variable type is already defined for a variable, `var name string = "jaime"`, then it can't be assigned to an `int` or `float` after. If you want to change the value of the variable, it has to be the same type as defined before
2. When using `:=` to define a variable, this cannot be used outside of a function. To use `:=`, it has to be in a function or else you'll get an error

### How to Define variables
There's 4 different ways to define variables:

In the first way, you're explicitly telling Go what type you want your variable. If you put a different type that what you said, like `int`, or `string` for example, then you'll get an error.

In the second way, Go is seeing what type you put in and automatically putting the variable type for you.

In the third way, it's pretty similar to the first way. The only difference is that the actual variable is being defined somewhere else in the code

In the fourth way, it's a short hand to the second method using `:=`. As stated before, these cannot be used outside a function

Below will be examples on the variable types and how they're defined

### Strings
```go
// first way
var nameOne string = "jaime"
// second way
var nameTwo = "eli"
// third way
var nameThree string
nameThree = "ivy"  // this could be defined later in the function

nameFour := "caca"
```

### Integers
```go
var ageOne int = 22  // first
var ageTwo = 21  // second
var ageThree int  // third
ageThree = 67
ageFour := 4  // fourth
```

### Bits and Memory
```go
var numOne int8 = 25
var numTwo int8 = -128
var numThree uint16 = 256
```
Something to note about these variable types is that they have maxes. `numTwo` is restricted to numbers from -128 to 128 since this is how bits work. And there's unsigned numbers as well that range from 0 to whatever.

### Floats
```go
var scoreOne float32 = 25.98
var scoreTwo float64 = 5474547467.67
scoreThree := 1.5  // scoreThree will be a float64
```
Floats also have the same thing with bits like we talked about before. Usually, `float64` are the default ones and will automatically be assigned to variables if not defined


## 4. Printing and Formatting Strings
Just for context, we'll be inside this function. Additionally, `fmt` is a standard library package that provides functions for formatted input and output.
```go
package main
import "fmt"

func main() {
    age := 35
    name := "jaime"
}
```
### Print
`fmt.Print()` will not give you a new line after printing something. `fmt.Println()` will though. Which is why if we ran this code, the output will be `HelloWorld!`. Additionally, it's simple to print out variables as well
```go
// Print
fmt.Print("Hello\n")

// Println
fmt.Println("World!")
fmt.Println("my age is", age, "and my name is", name)

//Printf (formatted sring) %_ = format specifier
fmt.Printf("my age is %v and my name is %v\n", age, name)
fmt.Printf("my age is %q and my name is %q\n", age, name)  // %q puts quotes around the string
fmt.Printf("age is of type %T\n", age)  // %T gives you the type of the variable
fmt.Printf("you scored %f points!\n", 225.55)  // %f is for floats and you're allowed to put 0.1, 0.01, etc to round the decimal points
fmt.Printf("you scored %0.1f points!\n", 225.55)

// Sprintf (save formatted atrings)
var str = fmt.Sprintf("my age is %v and my name is %\n", age, name)
fmt.Println("Saved string is:", str)
```
Using `fmt.Printf()` is a bit more complicated. It's very similar to how C handles printing variables. With the `%_` and a letter after (_ represents a letter)

`fmt.Sprintf()` allows you to store these formatted string into a variable.


## 5. Arrays and Slices
### Arrays (Static Arrays)
To define an array we can do this
```go
var ages [3]int = [3]int{20, 25, 22}
var ages = [3]int{20, 25, 22}  // this is the same as the one above

names := [4]string{"eli", "jaime", "jcast", "ivy"}
```
Since arrays are like static arrays, the size of the array is defined and cannot be altered. Meaning that you can't add elements to it that's greater than the specified size. Additionally, you cannot mix in integers, string, etc, together.

### Slices (Dynamic Arrays)
```go
var scores = []int{100, 50, 60}
scores[2] = 25
scores = append(scores, 85)

// slice ranges
rangeOne = := names[1:3]
rangeTwo := names[2:]
rangeThree := names[:3]

rangeOne = append(rangeOne, "bob")
```
Since slices are like dynamic arrays, you don't have to specify the size of the array and you can append as much elements as you want. 

Additionally, slice ranges are very similar to Python as the first element is inclusive but the last one is not


## 6. The Standard Library
Go has many packages available that do many useful things like sorting arrays, searching for values, etc. Here are a couple:
```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there !"

	fmt.Println(strings.Contains(greeting, "hello"))  // checks to see if the first parameter contains the second parameter and returns a boolean 
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))  // replaces the second parameter with the third parameter in the first parameter (this DOES NOT alter the original string, greeting in this case)
	fmt.Println(strings.ToUpper(greeting))  // changes the letters to all uppercase in the parameter (doesn't alter)
	fmt.Println(strings.Index(greeting, "ll"))  // finds the index of where the second parameter appears in the first parameter. if the second parameter doesn't exist in the first parameter, it returns the length + 1 of the first parameter
	fmt.Println(strings.Split(greeting, " "))  // splits the first parameter into an array depending on the second parameter

	ages := []int{45, 20, 78, 53, 25, 12, 56}
	sort.Ints(ages)  // sorts the array from smallest to largest (this DOES alter the parameter)
	fmt.Println((ages))

	index := sort.SearchInts(ages, 53)  // finds the index where the parameter appears
	fmt.Println(index)

	names := []string{"eli", "jaime", "jcast", "ivy"}
	sort.Strings(names)  // from smallest to largest again
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "eli"))  // finds the index where the parameter appears
}
```


## 7. Loops
There are a couple ways to do a loop in Go.

#### While Loop
This first one is basically a `while` loop. Go doesn't have a `while` command like most other programming lamguages but it's basically the same thing
```go
x := 0
for x < 5 {
    fmt.Println(x)
    x++
}
```

#### For Loop
Another way to do it is using a basic `for` loop. The syntax is very similar to Java
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

names := []string{"eli", "jaime", "jcast", "ivy"}
for i := 0; i < len(names); i++ {
    fmt.Println(names[i])
}
```

#### For ... Range Loop
Lastly, there's the `for ... range` loop. This allows you to get the index and value using the `range` function inside the for loop.
```go
for index, value := range names {
    fmt.Println(index, value)
}
```

If you don't want the index or the value, you can simply do this to not include them because if you don't use the values, you will get an error
```go
// no index
for _, value := range names {
    fmt.Println(value)
}

// no value
for index := range names {
    fmt.Println(index)
}
```


## 8. Booleans and Conditionals
### Booleans
Booleans in Go are the same as most programming languages. `false` and `true`

```go
age := 22
fmt.Println(age < 20)  // false
fmt.Println(age == 22)  // true
fmt.Println(age >= 29)  // false
fmt.Println(age != 33)  // true
```

### Conditionals
Also very similar to Java's syntax
```go
if age < 22 {
    fmt.Println("too low")
} else if age > 22 {
    fmt.Println("too high")
} else {
    fmt.Println("you got it")
}
```

This is for `if` statements inside of loops
```go
names := []string{"eli", "jaime", "jcast", "ivy"}
for index, value := range names {
    if index == 1 {
        fmt.Println("continuing at pos", index)
        continue
    }

    if index > 2 {
        fmt.Println("breaking at pos", index)
        break
    }

    fmt.Println(index, value)
}
```
What `continue` will do is if `index == 1` is `true`, then everything under it will not execute. It will just continue the for loop and skipping all the code under it. 

What `break` will do is if `index > 2` is `true`, then the for loop will break and it will no longer execute even if we didnt get through the entire names range