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