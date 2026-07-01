# Chapter 9. you're my type: Defined types
- You can use *any* type as an underlying type.
- How to define your *own* methods.

## Type errors in real life

## Defined types with underlying basic types
- If you have the following variable:

```go
var fuel float64 = 10
```

- You can use Go's defined types to make it clear what a vlue is to be used for.
  - Most commonly use structs as their underlying types.
  - They can be based on `int`, `float64`, `string`, `bool`, or any other type.

- Example:

```go
package main

import "fmt"

// Define two new types, each with an underlying type of float64
type Liters float64
type Gallons float64

func main() {
	var carFuel Gallons	// Define a variable with a type of Gallons.
	var busFuel Liters	// Define a variable with a type of Liters.
	carFuel = Gallons(10.0)	// Convert a float64 to Gallons.
	busFuel = Liters(240.0) // Convert a float64 to Liters.
	fmt.Println(carFuel, busFuel) // 10 240
}
```

- Once you have defined a type, you can do a conversion to that type from any value of the underlying type.
  - You write the type you want to convert to, followed by the value you want to convert in parentheses.

- Using short variable declaration:

```go
carFuel := Gallons(10.0)
busFuel := Liters(240.0)
```

- If you have a variable that uses a defined type
  - You *cannot* assign a value of a different defined type to it, even if the other type has the same underlying type.

```go
carFuel = Liters(240.0)	// cannot use Liters(240) (type Liters) as type
												// Gallons in assignment
busFuel = Gallons(10.0)	// cannot use Gallons(10) (type Gallons) as type
												// Liters in assignment
```

- You can convert between types that have the sume underlying type.
  - `Liters` can be converted to `Gallons` and vice versa.
  - Go only considers the value of the underlying type when doing a conversion.

```go
carFuel = Gallons(Liters(40.0))
busFuel = Liters(Gallons(63.0))
fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel) // Gallons: 40.0 Liters: 63.0
```

- You want to perform whatever operations are necessary to convert the underlying type value to a value appropriate for the type you are converting to.

```go
carFuel = Gallons(Liters(40.0) * 0.264)
busFuel = Liters(Gallons(63.0) * 3.785)
fmt.Printf("Gallons: %0.1f Liters: %0.1f\n", carFuel, busFuel) // Gallons: 10.6 Liters: 238.5
```

## Defined types and operators
- A defined type supports all the same oeprations as its underlying type.

```go
fmt.Println(Liters(1.2) + Liters(3.4))	// 4.6
fmt.Println(Gallons(5.5) - Gallons(2.2))	// 3.3
fmt.Println(Liters(2.2) / Liters(1.1))	// 2
fmt.Println(Gallons(1.2) == Gallons(1.2))	// true
fmt.Println(Liters(1.2) < Liters(3.4))	// true
fmt.Println(Liters(1.2) > Liters(3.4))	// false

type Title string
fmt.Println(Title("Alien") == title("Alien"))	// true
fmt.Println(Title("Alien") < Title("Zodiac"))	// true
fmt.Println(Title("Alien") > Title("Zodiac"))	// false
fmt.Println(Title("Alien") + "s")	// Aliens
fmt.Println(Title("Jaws 2" - " 2")) // invalid operation: Title("Jaws 2") - " 2" (operator - not defined on string)
```

- A defined type can be used in operations together with literal values

```go
fmt.Println(Liters(1.2) + 3.4)	// 4.6
fmt.Println(Gallons(5.5) - 2.2)	// 3.3
fmt.Println(Gallons(1.2) == 1.2)	// true
fmt.Println(Liters(1.2) < 3.4)	// true
```

- Defined types cannot be used in operations together with values of a different type, even if the other type has the same underlying type.

```go
fmt.Println(Liters(1.2) + Gallons(3.4))	// invalid operation: ...
fmt.Println(Gallons(1.2) == Liters(1.2))	// invalid operation: ...
```

## Converting between types using functions

## Fixing our function name conflict using methods

## Defining methods
- A method definition is very similar to a function definition.
  - There is really only one difference: you add one extra parameter, a **receiver parameter**, in parentheses before the function name.

```go
func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}
```

- To call the method

```go
value := MyType("a MyType value")
value.sayHi()
```

- The name of the receiver parameter in the method definition isn't important, but its type is.
  - The method you are defining becomes associated with all values of that type.

## The receiver parameters is (pretty much) just another parameter
- The type of the receiver parameter is the type that the method becomes associated with.
- Asides from that, the receiver parameter doesn't get special treatment from Go.
- Go lets you name a receiver parameter whatever you want, but it is more readable if all the methods you define for a type have receiver parameters with the same name.
  - By convention, Go developers usually use a name consisting of a single letter - the first letter of the receiver's type name, in lowercase.

>Go uses receiver parameters instead of the "self" or "this" values in other languages.

## A method is (pretty much) just like a function
- As with any other fuction, you can define additional parameters within parantheses following the method name.
  - These parameter variables can be accessed in the method block, along with the receiver parameter.
- As with any other function, you can declare one or more return values for a method, which will be returned when the method is called.
- As with any other function, a method is considered exported from the current package if its name begins with a capital letter, and it is considered unexported if its name begins with a lowercase letter.

## Pointer receiver parameters
- When you call a method that requires a pointer receiver on a variable with a nonpointer type, Go will automatically convert the receiver to a pointer for you.
- The same is true for variables with pointer types; if you call a method requiring a value receiver, Go will automatically get the value at the pointer for you and pass that to the method.

```go
// Package, imports omitted
type MyType string

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := MyType("a value")
	pointer := &value
	value.method()					// Method with value receiver
	value.pointerMethod()		// Method with pointer receiver
	pointer.method()				// Method with value receiver
	pointer.pointerMethod()	// Method with pointer receiver
}
```
- This code breaks convention.
- For consistency: all of your type's methods can take value receivers, or they can all take pointer receivers.
  - You should avoid mising the two.