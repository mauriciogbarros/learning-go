# Module 1: Functions and Organization
## 1 - Why use functions?
### What is a function
- A set of instructions with a name

```go
func main() {
	fmt.Printf("Hello, world.")
}
```

```go
func PrintHello() {
	fmt.Printff("Hello, world.")
}

func main() {
	PrintHello()
}
```

- Function declaration, name, call

### Reusability
- You only need to declare a function once
- Good for commonly used operations
- Graphics editing program might have `ThresholdImage()`
- Database program might have `QueryDbase()`
- Music program might have `ChangeKey()`

### Abstraction
- Details are hidden in the function
- Only need to understand operations
- Improves understandability

```go
func FindPupil() {
	GrabImage()
	FilterImage()
	FindEllipses()
}
```

---

## 2 - Function parameters and return values
### Function parameters
- Functions may need input data to perform their operations
- **Parameters** are listed in parenthesis after the function name
- **Arguments** are supplied in the call

```go
func foo(x int, y int) {
	fmt.Print(x * y)
}

func main() {
	foo(2, 3)
}
```

### Parameter options
- If no parameters are needed, put nothing in parentheses
- Still need parentheses

```go
func foo() { ... }
```

- List arguments of same type

```go
func foo(x, y int) { ... }
```

### Return values
- Functions can return a value as a result
- **Type of return value** after parameters in declaration
- Function call used on right-hand side of an assignment

```go
func foo(x int) int {
	return x + 1
}

y := foo(1)
```

### Multiple return values
- Multiple value types must be listed in declaration

```go
func foo2(x int) (int, int) {
	return x, x + 1
}

a, b := foo2(3)
```

---

## 3 - Call by value, reference
### Call by value
- Passed arguments are copied to parameters
- Modifying parameters has no effect outside the function

```go
func foo(y int) {
	y = y + 1
}

func main() {
	x := 2
	foo(x)
	fmt.Print(x)
}
```

### Tradeoffs of call by value
- Advantage: data encapsulation
  - Function variables only changed inside the function
- Disadvantage: copying time
  - Large objects may take a long time to copy

### Call by reference
- Programmer can **pass a pointer** as an argument
- Called function has direct access to caller variable in memory

```go
func foo(y *int) {
	*y = *y + 1
}

func main() {
	x := 2
	foo(&x)
	fmt.Print(x)
}
```

### Tradeoffs of call by reference
- Advantage: copying time
  - Don't need to copy arguments
- Disadvantage: data encapsulation
  - Function variables may be changed in called functions

---

## 4 - Passing arrays and slices
### Passing array arguments
- Array arguments are copied
- Arrays can be big, so this can be a problem

```go
func foo(x [3]int) int {
	return x[0]
}

func main() {
	a := [3]int{1, 2, 3}
	fmt.Print(foo(a))
}
```

### Passing array pointers
- Possible to pass array pointers

```go
func foo(x *[3]int) int {
	(*x)[0] = (*x)[0] + 1
}

func main() {
	a := [3]int{1, 2, 3}
	foo(&a)
	fmt.Print(a)
}
```

- Messy and unnecessary

### Pass slices instead
- **Slices contain a pointer** to the array
- Passing a slice copies the pointer

```go
func foo(sli int) int {
	sli[0] = sli[0] + 1
}

func main() {
	a := []int{1, 2, 3}
	foo(a)
	fmt.Print(a)
}
```

---

## 5 - Well-written functions
### Understandability
- Code is **functions** and **data**
- If you are asked to **find a feature**, you can find it quickly
  - "Where is the function that blurs the image?"
  - "Where do you compute the average score?"
- If you are asked about **where data is used**, you know
  - "Where do you modify the record list?"
  - "Where do you access the file?"

### Debugging principles
- Code crashes inside a function
- Two options for the cause
  - 1. Function is written incorrently
  - 2. Data that the function uses is incorrect

### Supporting debugging
- Functions need to be understandable
  - Determine if actual behavior matches desired behavior
- Data needs to be traceable
  - Where di that data come from?
  - Global variables complicate this

---

## 6 - Guidelines for functions
### Function naming
- Give functions a good name
  - Behavior can be understood at a glance
  - Parameter naming counts too

```go
func ProcessArray(a []int) float { ... }
func ComputeRMS(samples []float) float { ... }
```

- RMS = Root Mean Square
- `samples` is a slice of samples of a time-varying signal

### Functional cohesion
- Function should perform **only one operation**
- An "operation" depends on the context
- Example: Geometry application
- Good functions: `PointDist()`, `DrawCircle()`, `TriangleArea()`
- Merging behaviors make code complicated: `DrawCircle()` + `TriangleArea()`

### Few parameters
- Debugging requires tracing function input data
- More difficult with a large number of parameters
- Function may have functional cohesion
- `DrawCircle()` and `TriangleArea()` require different arguments

### Reducing parameter number
- May need to group related arguments into structures
- `TriangleArea()`, bad solution
  - 3 points, needed to define trianble
  - Each point has 3 floats (in 3D)
  - Total, 9 arguments
- `TriangleArea()`, goo solution
  - `type Point struct {x, y, z float}`
  - Total, 3 arguments

---

## 7 - Function guidelines
### Function complexity
- **Function length** is the most obvious measure
- Functions should be simple
  - Easier to debug
- Short functions can be complicated too

### Function length
- How do you write complicated code with simple functions?
- **Function call hierarchy**
- Option 1

```go
func a() {
	<100 lines>
}
```

- Option 2

```go
func a() {
	b()
	c()
}

func b() {
	<50 lines>
}

func c() {
	<50 lines>
}
```

### Control-flow complexity
- Control-flow describes conditional paths

```go
func foo() {
	if a == 1 {
		if b == 1 {
			...
		}
		...
	}
	...
}
```
- 3 control-flow paths

### Partitioning conditionals
- Functional hierarchy can reduce control-flow complexity

```go
func foo() {
	if a == 1 {
		CheckB()
	}
	...
}

func CheckB() {
	if b == 1 {
		...
	}
}
```