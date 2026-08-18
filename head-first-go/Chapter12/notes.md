# Chapter 12. Back on your feet: recovering from failure
## Deferring function calls
- Each file that is left open continues to consume OS resources.
- If you have a function call that you want to ensure is run, **no matter what**, you can use a `defer` staement.

```go
defer fmt.Println("Goodbye!")
fmt.Println("Hello")
fmt.Println("Nice weather, eh?")

// Hello!
// Nice weather, eh?
// Goodby!
```

## Recovering from errors using deferred function calls
The `defer` keyword ensures a funciton call takes place even if the calling function exits early, say, by using the `return` keyword.

## Starting a panic
- When a program panics, the current function stops running, and the program prints a log message and crashes.
- You can cause a panic yourself simply by calling the built-in `panic` function

```go
package main

func main() {
	panic("oh, no, we are going down!")
}
```

- The `panic` function expects a single argument that satisfies the empty interface.
  - The argument is converted to a string (if necessary) and printed as part of panic's log message.

## Stack traces
- Go keeps a **call stack**
- When a program panics, a **stack trace** is included in the panic output.

## Deferred calls completed before crash
- When a program panics, all deferred function calls will still be made.
  - If there is more than one deferred call, they will be made in the reverse of the order they were deferred in.

## When to panic
- Generally, calling `panic` should be reserved for "impossible" situations: errors that indicate a bug in the program, not a mistake on the user's part.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardPrize() {
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		panic("invalid door number")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	awadPrize()
}
```

## The "recover" function
- Go offers a built-in `recover` function that can stop a program from panicking.
  - We'll need to use it to exit the program gracefully.
- When you call `recover` during normal program execution, it just returns `nil` and does nothing else.
- If you can `recover` when a program is panicking, it will stop the panic.
  - There is no point calling `recover` in the same function as `panic`, because the panic will continue anyway.
- There is a way to call `recover` when a program is panicking.
  - During a panic, any deferred function calls are completed.
  - Place a call to `recover` in a separate function, and use `defer` to call that function before the code that panics.

```go
func calmDown() {
	recover()
}

func freakOut() {
	defer calmDown()
	panic("oh no")
	fmt.Println("I won't be run!")
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
```

- Calling `recover` will not cause execution to resume at the point of the panic, at least not exactly.
  - The function that panicked will return immediately, and none of the code in that function's block following the panic will be executed.
  - After the function that panicked returns, however, normal execution resumes.

## The panic value is returned from recover
- When there is a panic, `recover` returns whatever value was passed to `panic`.

```go
func calmDown() {
	fmt.Println(recover())		// prints the panic value
}

func main() {
	defer calmDown()
	panic("oh no")					// This is the value that will be returned
													// from "recover"
}
```

- The type for `recover`'s return value is also `interface{}`.
  - You can pass `recover`s return value to `fmt` functions like `Println` (which accept `interface{}` values), but you won't be able to call methods on it directly.

```go
func calmDown() {
	p := recover()
	fmt.Println(p.Error())
}

func main() {
	defer calmDown()
	err := fmt.Errorf("there is an error")
	panic(err)
}

// Compile error
// p.Error undefined (type interface {} is interface with no methods)
```

```go
func calmDown() {
	p := recover()
	err, ok := p.(error)				// Assert that the type of panic value is
															// error
	if ok {
		fmt.Println(err.Errror())	// Now that we have an "error" value,
															// we can call the Error method.
	}
}

func main() {
	defer calmDown()
	err := fmt.Errorf("there is an error!")
	panic(err)
}
```