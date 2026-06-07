# 3 - Call me

**Functions**

## Formatting output with Printf and Sprintf

- `Printf`: **print**, with **f**ormatting

`fmt.Printf("About one-third: %0.2f\n", 1.0 / 3.0)` => About one-third: 0.33

- `Sprintf`: like `Printf`, but it returns a a formatted string insted of printing it.

`resultString := fmt.Sprintf("About one-third: $0.2f\n", 1.0 / 3.0)`

### Formatting verbs (Printf & Sprintf)

- Any percent signs (%) are treated as the start of a **formatting verb**.
  - A section of the string that will be substituted with a value in a particular format.
- The remaining arguments are used as values with those verbs.

```go
fmt.Printf("The %s cost %d cents each.\n", "gumballs", 23)
fmt.Printf("That will be %f please.\n",  0.23 * 5)
```

- The letter following the percent sign indicates which verb to use.

| Verb  | Output                                                     |
| :---: | :--------------------------------------------------------- |
| `%f`  | Floating-point number                                      |
| `%d`  | Decimal integer                                            |
| `%s`  | String                                                     |
| `%t`  | Boolean (`true` or `false`)                                |
| `%v`  | Any value (chooses an appropriate format)                  |
| `%#v` | Any value, formatted as it would appear in Go program code |
| `%T`  | Type of the supplied value (`int`, `string`, etc.)         |
| `%%`  | A literal percent sign                                     |

- Make sure to add a newline (`\n`) at the end of each formatting string.
  - `Printf` does not automatically add a new line.

- `%#v` formatting verb can print values the way they would appear in Go code, rather than how they normally appear.
  - Can show some values that would otherwise be hidden.
  - `\t`, `\n`, ...

### Formatting value widths
- You can specify the minimum width after the percent sign for a formatting verb.
  - If the argument matching that verb is shorter than the minimum width, it will be padded with spaces until the minimum width is reached.

```go
fmt.Printf("%12s | %s\n", "Product", "Cost in cents")
fmt.Println("---------------------------")
fmt.Printf("%12s | %2d\n", "Stamps", 50)
fmt.Printf("%12s | %d\n", "Paper clips", 5)
fmt.Printf("%12s | %d\n", "Tape", 99)
```

**Output**
```
     Product | Cost in Cents 
-----------------------------
      Stamps | 50
 Paper Clips |  5
        Tape | 99
```

### Formatting fractional number widths
**%5.3f**
`%`: Start of the formatting verb
`5`: Minimum width of the entire number
`3`: Width after the decimal point
`f`: Formatting type

`fmt.Printf("%%7.3f: %7.f\n", 12.3456)` => %7.3f:   12.346
`fmt.Printf(" %%.2f: %.2f\n", 12.3456)` =>  %.2f: 12.35

## Declaring functions
```go
package main

import "fmt"

func sayHi() {
	fmt.Println("Hi!")
}

func main() {
	sayHi()
}
```

- Rules for functions names, same as the rules for variable names:
  - a name must begin with a letter, followed bby any number of additional letters and numbers.
  - Functions whose name begins with a capital letter are *exported*, and can be used outside the current package.
  - Names with multiple words should use `camelCase`.

### Declaring function parameters
```go
package main

import "fmt"

func main() {
	repeatLine("helo", 3)
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}
```

### Function return values
```go
package main

import "fmt"

func double(number float64) float64 {
	return number * 2
}

func main() {
	dozen : = double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))
}
```

### Error values
- An error value is any value with a method named `Error` that returns a string.
  - The simplest way to create one is to pass a string to the `errors` package's `New` function, which will return a new error value.
  - If you call the `Error` method on that error value, you will get the string you passed to `errors.New`.

```go
package main

import {
	"error"
	"fmt"
}

func main() {
	err := error.New("height can't be negative")
	fmt.Println(err.Error()) // height can't be negative
}
```

- If you are passing the error value to a function in the `fmt` or `log` package, you don't need to call its `Error` method.

- If you need to format numbers or other values in your error message:

```go
err := fmt.Errorf("a height of %0.2f is invalid.", -2.3333)
fmt.Println(err.Error())		// a height of -2.33 is invalid
fmt.Println(err)						// a height of -2.33 is invalid
```

### Declaring multiple return values
- To declare multiple return values for a function, place the return value types in a *second* set of parentheses in the function declaration, separated with commas.

```go
package main

import "fmt"

func manyReturns() (int, bool, string) {
	return 1, true, "hello"
}

func main() {
	myInt, myBool, myString := manyReturns()
	fmt.Println(myInt, myBool, myString)			// 1 true hello
}
```

- Names can be supplied for each return value

```go
package main

import (
	"fmt",
	"math"
)

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	cans, remainder := floatParts(1.26)
	fmt.Println(cans, remainder)				// 1 0.26
}
```

## Always handle errors!
- when a function returns an error value, it usually has to return a primary return value as well.
  - It's important to test whether the error value is `nil` before proceeding.

## The paintNeeded function
```go
package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height

	return area / 10.0, nil
}

func main() {
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err)		// Display error message and exits the program.
	}

fmt.Println("%0.2f liters needed\n", amount)	// 0.00 liters needed
}
```

## Pool Puzzle
```go
package main

import (
	"errors"
	"fmt"
)

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}

	return dividend / divisor, nil
}

func main() {
	quotient, err := divide(5.6, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}
}
```

## Function parameters receive copies of the arguments

## Pointers
- `&`: **address of** operator

```go
amount := 6
fmt.Println(amount)			// 6
fmt.Println(&amount)		// 0x1040a124
```

### Pointer types
- The type of a pointer is written with `*` symbol, followed by the type of the variable the pointer points to.
  - `*int` reads: pointer to `int`.
- We case `reflect.TypeOf` to show the types of pointers

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))		// Pointer to myInt
}
```

- We can declare variable that hold pointers.
  - A pointer variable can only hold pointers to one type of value.

```go
var myInt int
var myIntPointer *int
myIntPointer = &myInt
fmt.Println(myIntPointer)

// or
var myFloat float64
myFloatPointer := &myFloat
fmt.Println(myFloatPointer)
```

### Getting or changing the value at a pointer
- You can get the value of the variable a pointer refers to by typing the `*` operator right before the pointer in your code.

```go
myInt := 4
myIntPointer := &myInt
fmt.Println(myIntPointer)
fmt.Println(*myIntPointer)
```

- The `*` operator can also be used to update the value at a pointer:

```go
fmt.Println(myInt)
myIntPointer := &myInt
*myIntPointer = 8
fmt.Println(*myIntPointer)	// 8
fmt.Println(myInt)					// 8
```

## Code Magnets
```go
package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int
	myIntPointer = &myInt
	myInt = 42
	fmt.Println(*myIntPointer)
}
```

## Using pointers with functions
- It's possible to return pointers from functions.
  - In Go, it's OK to return a pointer to a variable that's local to a function.
  - Even though that variable is no longer in scope, as long as you still have the pointer, Go will ensure you can still access the value.

```go
func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func main() {
	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)
}
```

- You can also pass pointers to functions as arguments.

```go
func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func main() {
	var myBool bool = true
	printPointer(&myBoll)
}
```

### Fixing "double" function using pointers
```go
func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
}

func double(number *int) {
	*number *= 2
}
```
