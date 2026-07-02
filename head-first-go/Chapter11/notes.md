# Chapter 11. what can you do?: Interfaces
- Sometimes you don't care about the particular type of a value.
  - You just need to know that it will be able to do certain things.
- Go interfaces
  - They let you define variables and function parameters that will hold any type, as long as that type defines certain methods.

## Two different types that have the same methods

## Interfaces
- An interface is a set of methods that certain values are expected to have.
- Declaring an interface:

```go
type myInterface interface {
	methodWithoutParameters()
	methodWithParameter(float64)
	methodWithReturnValue() string
}
```

- Any type that has all the methods listed in an interface definition is said to **satisfy** that interface.
- A type that satisfies an interace can be used anywehre that interface is called for.
- A type can have methods in addition to those listed in the interface, but it mustn't be missing any, or it doesn't satisfy that interface.
- A type can satisfy multiple interfaces, and an interface can (and usually should) have multiple types that satisfy it.

## Defining a type that satisfies an interface
```go
package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
```

- Here is a quick program that will let us try `mypkg` out

```go
package main

import (
	"fmt"
	"mypkg"
)

func main() {
	var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParemeters()			// MethodWithoutParameters called
	value.MethodWithParameter(127.3)		// MethodWithParameter called with 127.3
	fmt.Println(value.MethodWithReturnValue())	// Hi from MethodWithReturnValue
}
```

## Concrete types, interface types
- A concrete type specifies not only what its values can do (what methods you can call on them), but also what they are: they specify the underlying type that holds the value's date.
- Interface types don't describe what a value is: they don't say what its underlying type is, or how its data is stored. They only describe what a value can do: what methods it has.

## Assign any type that satisfies the interface
```go
package main

import "fmt"

type Whistle string

funt (w Whistle) MakeSound() {
	fmt.Printon("Twee!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()								// Twee!
	toy = Horn("Toyco Blaster")
	toy.MakeSound()								// Honk!
}
```

- You can declare function parameters with interface types as well.

```go
func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	play(Whistle("Toyco Canary"))
	play(Horn("Toyco Blaster"))
}
```

## You can only call methods defined as part of the interface
- Once you assign a value to a variable (or method parameter) with an interface type, you can only call methods that are specified by the interface on it.

>If a type declares methods with pointer receivers, then you will only be able to use pointers to that type when assigning to interface variables.

```go
package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s == "off"
	} else {
		*s == "on"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func main() {
	s := Switch("off")
	// var t Toggleable = s
	var t Toggleable = &s
	t.toggle()
	t.toggle()
}
```

- When go decides whether a value satisfies an interface, pointer methods aren't included for direct values.
- But they are included for pointers.
- So the solution is to assign a pointer to a `Switch` to the `Toggleable` variable, instead of a direct `Switch` value

## Type assertions
- When you have a value of a concrete type assigned to a variable with an interface type, a type assertion lets you get the concrete type back.
  - It's *kind of* like a type conversion.

```go
var noiseMaker NoiseMaker = Robot("Botco Ambler")
var robot Robot = noiseMaker.(Robot)
```
- I know this variable uses the interface type `NoiseMaker`, but I am pretty sure *this* `NoiseMaker` is actually a `Robot`.
- Once you have used a type assertion to get a value of a concrete type back, you can call methods on it that are defined on that type, but aren't part of the interface.

## Type assertion failures
- Everything compiles successfully, but when we try to run it, we get a runtime panic!
  - Trying to assert that a `TapePlayer` is actually a `TapeRecorder` did not go well.

## Avoiding panics when type assertions fail
- If type assertions are used in a context where multiple return values are expected, they a have a second, optional value that indicates whether the assertion was successful or not.
  - A `bool`, and it will be `true` if the value's original type was the asserted type, or `false` if not.
  - It is usually assigned to a variable `ok`

```go
var player Player = gadget.TapePlayer{}
recorder, ok := player.(gadget.TapeRecorder)
if ok {
	recorder.Record()
} else {
	fmt.Println("Player was not a TapeRecorder")
}
```

## Testing TapePlayers and TapeRecorders using type assertions

## The "error" interface
- The `error` type is just an interface!
- It looks something like this:

```go
type error interface {
	Error() string
}
```

- Declaring the `error` type as an interface means that if it has an `Error` method that returns a `string`, it satisfies the `error` interface, and it's an `error` value.

```go
type ComedyError string								// Define a type with an
																			// underlying type of "string"

func (c ComedyError) Error() string {	// Satisfy the error interface

	return string(c)										// The Error method needs to
																			// to return a string, so do a
																			// type conversion
}

func main() {
	var err error
	err = ComedyError("What's a programmer's favorite beer? Logger!")
	fmt.Println(err)
}
```

- If you need an `error` value, but also need to track more information about the error than just an error message string.
  - You can create your own type that satisfies the `error` interface **and** stores the information you want.

```go
type OverheatError float64
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	var err error = checkTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
```

## The Stringer interface
- The `fmt` package defines the `fmt.Stringer` interface
  - Allows any type to decide how it will be displayed when printed.
  - Just define a `String()` method that returns a `string`.

```go
type Stringer interface {
	String() string
}
```

- Example

```go
type CoffeePot string
func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	coffeePot := CofeePot("LuxBrew")
	fmt.Println(coffeePot.String())		// LuxBrew coffee pot
}
```

- Many functions in the `fmt` package check whether the values passed to them satisfy the `Stringer` interface, and call their `String` methods if so.
  - This includes `Print`, `Println`, and `Printf` functions and more.

## The empty interface
- What would happen if we declared an interface type that didn't require any methods at all?
  - It would be satisfied by any type => It would be satisfied by all types.

```go
type Anything interface { }
```

- The type `interface{}` is know as **the empty interface**.
- It is used to accept values of any type.
- Example: Println
  - go doc: func Println(a ...interface{}) (n int, err error)
- Remember, you can only call methods on it that are part of the interface.
  - That means there are not methods you can call ona value with the empty interface type.
  - To call methods on a value with the empty interface type, you would need to use a type assertion to get a value of the concrete type back.
  - By this point, you are probably better off writing a function that accepts only that specific concrete type.