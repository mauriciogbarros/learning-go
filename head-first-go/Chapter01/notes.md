# 1 - Syntax Basics
## What does it all mean?
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

- Every Go file starts with a `package` clause.
  - Pakcage `main` is required if the code is going to be run directly.
  - `main` function is special, it gets run first when the program runs.
- The typical Go file layout
  - 1. The package clause
  - 2. Any `import` statements
  - 3. The actual code
- No semicolons
- Go developrs expect the standard Go format.
  - Indentation and spacing
  - Run `go fmt` to automatically fix formatting.

## What if something goes wrong?
- Fail reasons
  - Every Go file has to begin with a package clause.
  - Every go file has to import every package it references.
  - Go files must import **only** the pakcages they reference.
  - Go looks for a function named `main` to run first.
  - Everything in Go is case-sensitive.
  - The package name is required before the function call.

## Using functions from other packages
- To import multiple packages: `import (...)`

```go
package main

import (
	"fmt"
	"math"
	"string"
)

func main() {
	fmt.Println(math.Floor(2.75))
	fmt.Println(string.Title("head first go"))
}
```

## Strings
|Escape sequence|Value|
|:---:|:---:|
|`\n`|A new line character.|
|`\t`|A tab character.|
|`\"`|Double quotation marks.|
|`\\`|A backlash.|

## Runes
- Used to represent single characters.
- Written with single quotes: `'A'`
- Escpae sequences can be used in a rune literal.
- Go uses the Unicode standard for storing runes.

## Booleans
- `true` or `false`

## Numbers
- Go treats integer and floating-point numbers as different types.
- A decimal point can be used to distinguish an integer from a floating-point number.

## Types
- Go is **statically typed**
- View type of any value: import "reflect", call reflect.TypeOf()
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(42))						// int
	fmt.Println(reflect.TypeOf(3.1415))				// float64
	fmt.Println(reflect.TypeOf(true))					// bool
	fmt.Println(reflect.TypeOf("Hello, Go!"))	// string
}
```

## Declaring variables
```go
var quantity int
var length, width float64
var customerName string

length, width = 1.2, 2.4
```

```go
package main

import "fmt"

func main() {
	var quantity int
	var length, width, float64
	var customerName string

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length * width, "square meters")
}
```

## Zero values
- If you declare a variable without assigning it a value, that variable will contain the **zero value** for its type.
  - For numeric types: 0
  - For string: empty string
  - for boolean: `false`

## Short variable declarations
- Declaring variables and assigning values on the same line:
```go
var quantity int = 4
var length, width float64 = 1.2, 2.4
var customerName string = "Damon Cole"
```

- If you know the initial value of a variable:
  - `:=` assignment operator
```go
quantity := 4
length, width := 1.2, 2.4
customerName := "Damon Cole"
```

- Breaking stuff
  - A variable can only be declared once.
  - Variables can only be assigned values of the same type
  - A value is required for every variable being assigned.
  - All declared variables must be used.

## Naming rules
- Enforced by the language:
  - A name must begin with a letter
  - If the name of a variable, function, or type begins with a capital letter, it is considered **exported**
  - If the name of a variable, function, or type begins with a lowercase letter, it is considered **unexported**
- Conventions
  - Use *camel case*

## Conversions
- Math and comparison operations in Go require that the included values be of the same type.

```go
var myInt int = 2
float64(myInt)
```
- Converting from `float64` to `int` truncates.

## Compiling Go code
>Source code (hello.go) => Compiler => Compiled code (executable file)

## Go tools
|Command|Description|
|:---:|:---|
|`go build`|Compiles source code files into binary files.|
|`go run`|Compiles and runs a program, without saving an executable file.|
|`go fmt`|Reformats source files using Go standard formatting.|
|`go version`|Displays the current Go version.|