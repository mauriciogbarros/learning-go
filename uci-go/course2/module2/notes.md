# Module 2: Function types
## 1 - First-class values
### Functions are first-class
- Functions can be treated like other types
  - Variables can be declared with a function type
  - Can be created dynamically
  - Can be passed as arguments and returned as values
  - Can be stored in data structures

### Variables as functions
- Declare a variable as a func

```go
var funcVar func(int) int	func incFn(x int) int {
	return x + 1
}

func main() {
	funcVar = incFn
	fmt.Print(funcVar(1)) // 2
}
```

- Function on right-hand side, without `()`

### Functions as arguments
- Function can be passed to another function as an argument

```go
func applyIt(afunct func(int) int, val int) int {
	return afunc(val)
}

func incFn(x int) int { return x + 1 }
func decFn(x int) int { return x - 1 }

func main() {
	fmt.Println(applyIt(incFn, 2)) // 3
	fmt.Println(applyIt(decFn, 2)) // 1
}
```

### Anonymous functions
- Don't need to name a function

```go
func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func main() {
	v := applyIt(func (x int) int { return x + 1}, 2)
	fmt.Println(v)
}
```

## 2 - Returning functions
### Function as return values
- Functions can return functions
- Might create a function with controllable parameters
- Example: Distance to origin function
  - Takes a point(x, y coordinates)
  - Returns distance to origin
- What if I want to change the origin?
  - Option 1: Pass origin as argument
  - Option 2: Define function with new origin

### Function defines a function
```go
func MakeDistOrigin (o_x, o_y float64) func(float64, float64) float64 {
	fn := func (x, y float64) float64 {
		return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
	}

	return fn
}
```
- Origin location is passed as an argument
- Origin is built into the returned function

### Special-purpose functions
```go
func main() {
	Dist1 := MakeDistOrigin(0, 0)
	Dist2 := MakeDistOrigin(2, 2)
	fmt.Println(Dist1(2, 2))	// 2.82...
	fmt.Println(Dist2(2, 2))	// 0
}
```
- `Dist1()` and `Dist2()` have different origins

### Environment of a function
- Set of all names that are valid inside a function
- Names defined locally, in the function
- Lexical scoping
- Environment includes names defined in block where the function is defined.

```go
var x int
func foo(y int) {
	z := 1
}
```

### Closure
- Function + its environment
- When functions are passed/returned, their environment comes with them.

```go
func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func (x, y float64) float64 {
		return math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2))
	}
}
```
- `o_x` and `o_y` are in the flosure of `fn()`

## 3 - Variadic and Deferred
### Variable argument number
- Functions can take a variable number of arguments
- Use elipsis `...` to specify
- Treated as a slice inside function

```go
func getMax(vals ...int) int {
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}
```

### Variadic slice argument
```go
func main() {
	fmt.Println(getMax(1, 3, 6, 4))
	vslice := []{1, 3, 6, 4}
	fmt.Println(getMax(vslice...))
}
```
- Can pass a slice to a variadic function
- Need the `...` suffix

### Deferred function calls
- Call can be **deferred** until the surrounding function completes
- Typically used for cleanup activities

```go
func main() {
	defer fmt.Println("Bye!")
	fmt.Println("Hello!") // Hello!
												// Bye!
}
```

### Deferred call arguments
- Arguments of a deferred call are evaluated immediately

```go
func main() {
	i := 1
	defer fmt.Println(i + 1)
	i++
	fmt.Println("Hello!") // Hello!
												// 2
}
```

