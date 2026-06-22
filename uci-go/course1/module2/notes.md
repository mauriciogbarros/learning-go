# Module 2 - Basic Data Types
## 1. Pointers
- A pointer is an address to data in memory
- `&` operator returns the address of a variable/function
- `*` operator returns data at an address (dereferencing)
- Example:

```go
var x int = 1
var y int
var ip *int // ip is point to int

it = &x			// ip now points to x

y = *ip			// y is now 1
```

### New
- Alternate way to create a variable
- `new()` function creates a variable and returns a pointer to the variable.
- Variable is initialized to zero

```go
ptr += new(int)
*ptr = 3
```

## 2. Variable Scope
- The places in code where a variable can be accessed

### Blocks
- A sequence of declarations and statements within matching brackets, `{}`
  - Including function definitions
- Hierarchy of **implicit blocks** also
  - Universe block - all Go source
  - Package block - all source in a package
  - File block - all source in a file
  - "if", "for", "switch" - all code inside the statement
  - Clause in "switch" or "select" - individual clauses each get a block

### Lexical scoping
- Go is a lexically scoped using blocks
  - b_i >= b_j if b_j is defined inside b_i
  - "defined inside" is transitive

```go
var x = 4									// b1
func f() {
	fmt.Printf("%d", x)			// b2
}
func g() {
	fmt.Printf("%d", x)			// b3
}
```

### Scope of Variables
- Variable accessible from b_j if:
  - 1) Variable is declared in block b_i, and
  - 2) b_i >= b_j

## 3. Deallocating memory
### Deallocating space
- When a variable is no longer needed, it should be **deallocated**
  - Memory space is made available
- Otherwise, we will eventually run out of memory

```go
func f() {
	var x = 4
	fmt.Printf("%d%, x)
}
```
- Each call f() creates an integer

### Stack vs. Heap
#### General concept
- Stack is dedicated to function calls
  - Local variables are stored here
  - Deallocated after function completes
- Heap is a persistent region of memory
  - Needs to be explicitly deallocated

### Manual Deallocation
- Data on the heap must be deallocated when it is done being used
- In most compiled languages (i.e. C), this is done manually
```c
x = malloc(32);
free(x);
```
- Error-prone, but fast

## 4. Garbage Collection
### Pointers and Deallocation
- Hard to determine when a variable is no longer in use

```go
func foo() *int {
	x := 1
	return &x
}

func main() {
	var y *int
	y = foo()							// foo returns pointer to x
												// main can still use that address
	fmt.Printf("%d", *y)
}
```

### Garbage Collection
- In interpreted languages, this is done by the **interpreter**
  - Java Virtual Machine
  - Python Interpreter
- Easy for the programmer
- Slow (need an interpreter)

### Garbage Collection in Go
- Go is a compiled language which enables garbage collection
- Implementation is fast
- Compiler determines stack vs heap
- Garbage collection in the background

## 5. Comments, Printing, Integers
### Comments
- Comments are text for understandability
- Ignored by the compiler
- Single-line comments `//`
- Block comments `/* ... */`

### Printing
- Import from the fmt package
- `fmt.Printf()` (fmt.Println) prints a string

```go
fmt.Printf("Hi")
x := "Joe"
fmt.Printf("Hi " + x)
```

- Format strings are good for formatting
  - Conversion characters for each argument
  - `fmt.Printf("Hi %s", x)`

### Integers
- Generic int declaration: `var x int`
- Different lengths and signs
  - int8, int6, int32, int64, uint8, uint16, uint32, uint64
- Binary operators
  - Arithmetic: +, -, *, /, %, <<, >>
  - Comparison: ==, !=, >, <, >=, <=
  - Boolean: &&, ||

## 6. Ints, Float, Strings
### Type Conversions
- Most binary operations need operands of the same type
  - Including assignments

```go
var x int32 = 1
var y int16 = 2
x = y							// Does not work
```
- Convert type with `T()` operation
  - `x = int32(y)`

### Floating point
- `float32` ~6 digits of precision
- `float64` ~15 digits of precision
- Expressed using decimals or scientific notation

```go
var x float64 = 123.45
var y float64 = 1.2345e2
```

- Complex numbers represented as two floats: real and imaginary
  - `var z complex128 = complex(2, 3)` // 2 + 3i

### ASCII and Unicode
- American Standard Code for Information Exchange
- Character coding - each character is associated with a (7) 8-bit number
  - 'A' = 0x41
- **Unicode** is a 32-bit character code
- **UTF-8** is variable length
  - 8-bit UTF codes are same as ASCII
  - Default in Go
- **Code points** - Unicode characters
- **Rune** - a code point in Go

### Strings
- Sequence of arbitrary bytes
  - Read-only
  - Often meant to be printed
- **String literal** notated by double quotes
  - `x := "Hi there"`
- Each by is a rune (UTF-8 code point)

## 7. Strings Packages
- Unicode Package
  - Runes are divided into many different categories
  - Provides a set of functions to test categories of runes
    - `IsDigit(r rune)`
    - `IsSpace(r rune)`
    - `IsLetter(r rune)`
    - `IsLower(r rune)`
    - `IsPunct(r rune)`
  - Some functions perform conversions
    - `ToUpper(r rune)`
    - `ToLower(r rune)`
- Strings Package
  - Functions to manipulate UTF-8 encoded strings
  - String search functions
    - `Compare(a, b)` - returns an integer comparing two string lexicographically. 0 if a == b, -1 a < b, and +1 a > b.
    - `Contains(s, substr)` - returns true if substring is inside s
    - `HasPrefix(s, prefix)` - returns true if the string s beigns with prefix
    - `Index(s, substr)` - reeturns the index of the first instance of substring s
  - String Manipulation
    - Strings are immutable, but modified strings are returned
    - `Replace(s, old, new, n)` - replace returns a copy of the string s with the first n instances of old replaced by new
    - `ToLower(s)`
    - `ToUpper(s)`
    - `TrimSpace(s)` - returns a new string with all leading and trailing white space removed
- Strconv Package
  - Conversion to and from string representations of basic data types
  - `Atoi(s)` - converts strings s to int
  - `Itoa(s)` - converts int (base 10) to string
  - `FormatFloat(f, fmt, prec, bitSize)` - converts floating point number to a string
  - `ParseFloat(s, bitSize)` - converts a string to a floating point number

## 8. Constants
- Expression whose value is known at compile time
- Type is inferred from righthand side (boolean, string, number)

```go
const x = 1.3
const (
	y = 4
	z = "Hi"
)
```

## iota
- Generate a set of related but distinct constants
- Often represents a property which has several distinct possible values
  - Days of the week
  - Months of the year
- Constants must be different but **actual value is not important**.
- Like an enumerated type in other languages

```go
type Grades int
const (
	A Grades = iota
	B
	C
	D
	F
)
```
- Each constant is assigned to a unique integer
  - Starts at 1 and increments

## 9. Control Flow
### Control Structures
- Statements which alter control flow
  - `if <condition> { <consequent> }`
    - Expression `<condition>` is evaluated
    - `<consequent>` statements are executed if condition is `true`
```go
if x > 5 {
	fmt.Printf("Yup")
}
```
  - For loops
    - Iterates while a condition is true
    - May have an initialization and update operation
    - `for <init>; <cond>; <update> { <stmts> }`

```go
for i := 0; i < 10; i++ {
	fmt.Printf("hi ")
}

i = 0
for i < 10 {
	fmt.Printf("hi ")
	i++
}

for {
	fmt.Printf("hi ")
}
```

	- Switch/Case
  	- **Switch** is a multi-way if statement
  	- Switch may contain a **tag** which is a variable to be checked.
  	- Tag is compared to a constant defined in each **case**.
  	- Case which matches tag is executed

```go
switch x {
	case 1:
		fmt.Printf("case1")
	case 2:
		fmt.Printf("case2")
	default:
		fmt.Printf("nocase")
}
```

## 10. Control Flow, Scan
### Tagless Switch
- Switch may not contain a tag
- Case contains a boolean expression to evaluate
- First true case is executed

```go
switch {
	case x > 1:
		fmt.Printf("case1")
	case x < -1:
		fmt.Printf("case2")
	default:
		fmt.Printf("nocase")
}
```

### Break and Contiue
- **Break** exits the containing loop

```go
i := 0
for i < 10 {
	i++
	if i == 5 { break }
	fmt.Printf("hi ")
}
```

- **Continue** skips the rest of the current iteration

```go
i := 0
for i < 10 {
	i++
	if i == 5 { continue }
	fmt.Printf("hi ")
}
```

### Scan
- Scan reads user input
- Takes a pointer as an argument
- Typed data is written to pointer
- Returns number of scanned items

```go
var appleNum int
fmt.Printf("Number of apples?")
num, err := fmt.Scan(&appleNum)
fmt.Printf(appleNum)
```
