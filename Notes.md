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


## 9. Using Functions
A function can be defines as depicted below. The `func` keyword, the name of the function, and the parameters and parameter types
```go
func sayGreeting(n string) {
    fmt.Println("Good Morning", n)
}
```
So once this is passed through `main()`, it'll just say good morning to whatever name passed through

Another way to use funtions is by passing in more interesing parameters like slices (dynamic arrays) and another function
```go
func cycleNames(n []string, f func(string)) {
    for _, v := range n {
        f(v)
    }
}
func main() {
    cycleNames([]string{"jaime", "eli"}, sayGreeting)
}
```
So we're passing in an array of names and a function that takes in a string and that function does something to each array value inside the `for range` loop. Then in `main()`, we pass in an array and the `sayGreeting` function that we defined earlier. So this would just say good morning to each name in the array passed in

Another thing to note is that these functions are simply printing. If we wanted to `return` something, then we'd have to explicitly put that on the `func` lined as shown below
```go
func circleArea(r float64) float64 {
    return math.Pi * r * r
}
```


## 10. Multiple Return Values
As we saw before, if you want to return a value, you have to say what type you're going to return. If you want to return multiple values, then you'll just put it in paranthesis and say the types in there like shown in this example
```go
import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	i, j := getInitials("jaime castaneda")
	fmt.Println(i, j)

	k, l := getInitials("eli")
	fmt.Println(k, l)
}
```

## 11. Package Scopes
If you have two Go files, `main.go` and `greetings.go` for this example, and they're in the same directory, as long as they both have `package main` on top, then the global functions and variables defined in each one will work in the other. 

If in `greetings.go` I have this defined:
```go
var points = []int{25, 70, 56, 80}

func sayHello(n string) {
	fmt.Println("Hello", n)
}
```
It's fine for me to call them in `main.go` and everything will work fine and vice versa if I had something in `main.go` and wanted to call it in `greetings.go`:
```go
func main() {
	sayHello("mario")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}
```


## 12. Maps
Maps are very similar to dictionaries in Python. The main difference is that the key and value pairs have to remain the same type as defined. If you defined the key as a string and the value as an integer, you can't pass another type into them
```go
menu := map[string]float64{
		"soup":  4.99,
		"pie":   7.98,
		"water": 3.47,
	}
fmt.Println(menu)
fmt.Println(menu["pie"])  // to get just the value
```

This is how to loop through maps:
```go
for k, v := range menu {
    fmt.Println(k, "-", v)
}
```

To **update** an item in the map, you can do so like this:
```go
menu["water"] = 4.99
```


## 13. Pass by Value
### Go is a Pass-by-value language
* Go makes "copies" of values when passed into functions

If we have this code (this applies to strings, ints, bools, floats, arrays, and structs):
```go
func updateName(x string) {
	x = "wedge"
}

func main() {
	name := "jaime"
	updateName(name)
	fmt.Println(name)
}
```
You'd expect the result to be `wedge`. This is not the case though. `jaime` is still returned. Why does this happen? Go passes in a **copy** of the `name` variable to the `updateName` function. So it's not the original `name` variable. Just a copy of it. So the copy is being changed, not the original variable.

One way to fix this is by returning the value instead like the following:
```go
func updateName(x string) string {
	x = "wedge"
	return x
}

func main() {
	name := "jaime"
	name = updateName(name)
	fmt.Println(name)
}
```
If you do not include the `name = updateName(name)`, then the code will still return `jaime` as seen before. 

However, these types act differently compared to the ones listed above (slices, maps, functions)

```go
func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}
func main() {
    menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}
    updateMenu(menu)
	fmt.Println(menu)
}
```
In this instance, the coffee item will be added to this even though a copy was created. Why does this happen? `menu` is being passed in and this `menu` variable gets copied. But this copies variable still points to the same address aas the key value pairs in the `menu` variable. And since they're pointing to the same pairs, the `menu` variable will be updated.


## 14. Pointers
Pointers, just like in C, are a way to reference an address in memory. If we had a variable defined and we wanted to get the memory address of it, we can do so using the `&`. This will give us a memory address of `0x14000010040` or whatever it is for you.
```go
name := "jaime"
fmt.Println(&name)
```

To get the value at the specified memory address, we can use `*`
```go
m := name
fmt.Println(*m)
```
And with this we'd get `jaime`.

Using pointers allows us to fix the issue that we had before with Go passing in copies inside of funtions. Since now we'd be able to pass in the address and the copy will point to the same address as the variable and it'll update it
```go
func updateName(n *string) {
	*n = "wedge"
}

func main() {
	name := "jaime"
	m := &name

	updateName(m) // using pointer as arg, can pass &name
	fmt.Println(name)
}
```
And with this, we'd get `wedge` now.


## 15. Structs and Custom Types
Unlike most programming languages, Go does not have classes. Instead, Go uses custom types to create a `struct`. A `struct` is a blueprint a type of data
```go
type bill struct {
    name string
    items map[string]float64
    tip float64
}
```

And we can define a function to create a new bill
```go
func newBill(name string) bill {
    b := bill{
        name: name,
        items: map[string]float64{},
        tip: 0,
    }
    return b
}
```


## 16. Receiver Functions
If you want to call a function using dot notation, i.e. `myBill.format()`, then you'll need to build **receiver functions**. These can be defined in the same file as the struct was. To do so, you will need to pass in the bill struct **before** the name of the function and **after** the `func` keyword
```go
func (b bill) format() string {
    fs := "Bill Breakdown: \n"
    var total float64 = 0

    for k, v := range b.items {
        fs += fmt.Sprintf("%v ...$%v \n", k+":", v)
        total += v
    }
    fs += fmt.Sprintf("%v ...$%0.2f", "total:", total)
    return fs
}
```
So we're adding `(b bill)` after the `func` keyword. In this `format()` function, we're printing out the list of items and the total. So we initialize the variables, then we loop through `b`'s items (since every bill has items in it) and adding the value of the items to the total and then once that loop is over, adding the total to the string and printing it out


## 17. Receiver Functions with Pointers
We can create more receiver functions that can update and add more items to the bill list
```go
// update tip
func (b bill) updateTip(tip float64) {
    b.tip = tip
}

// add items to the bill
func (b bill) addItem(name string, price float64) {
    b.items[name] = price
}
```
But if you try to print these updates out in `main.go`, they won't actually appear. And the reason why is because of what we talked about before where Go creates a copy of the item passed in. It does the exact same thing here. 

To fix this, we can pass in pointers into the function like so
```go
func (b *bill) updateTip(tip float64) {
    b.tip = tip
}
```
Now it won't make a copy of the item and it'll work. It's also important to note that it's best practice to pass in pointers as much as you can so that you're not making so many copies of objects


## 18. User Input
To actually be able to read input from a user we need a reader. The `bufio` package provides us this and we can define it using `bufio.NewReader`. In this reader, however, we need to be able to get inpur from the user in the console. So, we can use the `os` package and define it using `os.Stdin` (standard input).
```go
func createBill() bill {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Create a new bill name: ")
    name, _ := reader.ReadString('\n')  // when the user clicks enter, the string will be read
    name = strings.TrimSpace(name)

    b := newBill(name)
    fmt.Println("Created the bill - ", b.name)

    return b
}
```
The function `createBill()` will return a `bill` type object. It will initalize a reader, print to the console the prompt, get the name, trim the whitespace in the name, initalize a varaible `b` using the name, print out the bill created with the name, and return the bill `b`

The only down side to this is that this can get repetative if we have multiple prompts that we want to get the user. So what we can do is create a different functions that gets the input instead
```go
func getInput(prompt string, r *bufio.Reader) (string, error) {
    fmt.Print(prompt)

    input, err := r.ReadString('\n')

    return strings.TrimSpace(input), err
}

// now createBill() can look like this
func createBill() bill {
    reader := bufio.NewReader(os.Stdin)
    name, _ := getInput("Create a new bill name: ", reader)

    b := newBill(name)
    fmt.Println("Created the bill - ", b.name)
    return b
}
```
In `getInput()`, we pass in the `prompt` as a `string` and `r` as a pointer to `bufio.Reader`. The reason we pass r as a pointer to `bufio.Reader` is because when we hover over the `reader` variable in `createBill()`, it says `var reader *bufio.Reader` so we have to pass that type in. Then we do exactly what we did earlier in getting the standard input put now using the parameters passed in.

Then in `createBill()` we just use this function now to make our lives easier. 

We can create a new function that gives the user some options to what they can do once their bill is created
```go
func promptOptions(b bill) {
    reader := bufio.NewReader(os.Stdin)
    opt, _ := getInput("Choose option (a - add an item, s - save the bill, t - add tip):", reader)
    fmt.Println(opt)
}
```
For now, this function doesn't work how we intend for it to work but it will once we finish the next couple of topics


## 19. Switch Statements
Switch statements are pretty straightforward and similar to other programming languages. You use the `switch` keyword with the parameter next to it, then you define the cases where this parameter is equal to the cases
```go
func promptOptions(b bill) {
    reader := bufio.NewReader(os.Stdin)
    opt, _ := getInput("Choose option (a - add an item, s - save the bill, t - add tip): ", reader)

    switch opt {
        // if it's a, we get the name of the item they want to add and the price and we print out the name and price
        case "a":
            name, _ := getInput("Item name: ", reader)
            price, _ := getInput("Item price: ", reader)
            fmt.Println(name, price)

        // if its t, we get the tip amount and print out the tip
        case "t":
            tip, _ := getInput("Enter tip amount ($): ", reader)
            fmt.Println(tip)

        // leaving this empty for now
        case "s":
            fmt.Println("You chose s")

        // if none of the other options were chose, then this gets printed out
        default:
            fmt.Println("That was not a valid option. Try again")
            promptOptions(b)  // so the user has to chose an actual option
    }
}
```
One important thing to note is that the name, price, and tip that the user types in will be a **string**. And as we saw before, the price and tip have to be a `float64`


## 20. Parsing Floats
To parse strings into floats we can use the `strconv` package and use it like the following
```go
case "a":
    name, _ := getInput("Item name: ", reader)
    price, _ := getInput("Item price: ", reader)

    p, err := strconv.ParseFloat(price, 64)  // 64 is for the bit size

    // if there is an error like a string 'h' for example being passed in, then this part executes
    if err != nil {
        fmt.Println("The price must be a number")
        promptOptions(b)
    }
    b.addItem(name, p)
    fmt.Println("Item added - ", name, price)
    promptOptions(b)
```
So we use `strconv.ParseFloat()` to get the string converted to a float and an error. Then we check if there is an error and if there is to try again, then we use the receiver functions defined before to add the item into the bill. This will the full code for the tip part
```go
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add an item, s - save the bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Println("Tip added - ", tip)
		promptOptions(b)
	case "s":
		fmt.Println("You chose to save the bill")

	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}
```


## 21. Saving Files
To save the bill, we'll create another receiver function
```go
func (b *bill) save() {
    data := []byte(b.format())

    err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

    if err != nil {
        panic(err)
    }
    fmt.Println("bill was saved to file")
}
```
So we pass in a pointer like we have been doing with the other receiver functions. The reason why our data is in bytes is because to write the function to a file, it has to be in bytes. So we do so like shown above. Then we assign `err` to the `os.WriteFile()` function. `"bills/"` comes from the directory name, `b.name` adds the name to the file, `".txt"` adds the extention that this will be a text file, `data` gives the data that we got above, and `0644` is for the permission.

If there is an error, then we use the `panic()` function that will stop everything. Then we print to the console that the bill was added. 

Then we add this to the `case "s":` and it will save the files to a bills folder


## 22. Interfaces
So imagine we have two structs like the following
```go
type square struct {
    length float64
}

type circle struct {
    radius float64
}
```

Then we define some receiver functions for them
```go
// square methods
func (s square) area() float64 {
    return s.length * s.length
}
func (s square) circumf() float64 {
    return s.length * 4
}

// circle methods
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
    return 2 * math.Pi * c.radius
}
```

And with these methods, we wanted to print out the shape
```go
func printShapeInfo(s shape) {
    fmt.Printf("area of %T is %0.2f \n", s, s.area())
    fmt.Printf("circumference of %T is %0.2f \n", s, s.circumf())
}
```
But if we want to pass in a circle or square, we'd only be able to pass in one of these. So if we want to be able to pass in both, we can create an interface
```go
type shape interface {
    area() float64
    circumf() float64
}
```
Now this will work for both shapes and we can pass this into main
```go
func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}
```