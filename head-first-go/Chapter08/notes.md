# Chapter 8. building storage: Structs
## Slices and maps hold values of ONE type
- Arrays, slices, and maps are no help if you need to mix values of different types.
  - They can only be set up to hold values of a single type.

## Structs are built out of values of MANY types
- A **struct** (short for "structure") is a vlue that is constructed out of other values of many different types.
- Declaration:

```go
var myStruct struct {
	number float64
	word string
	toggle bool
}
fmt.Printf("%#v\n", myStruct) // struct { number float64; word string; toggle bool } {number:0, word:"", toggle:false}
```

## Access struct fields using the dot operator
```go
var myStruct struct {
	number float64
	word string
	toggle bool
}
myStruct.number = 3.14
myStruct.word = "pie"
myStruct.toggle = true
fmt.Println(myStruct.number) // 3.14
fmt.Println(myStruct.word) // pie
fmt.Println(myStruct.toggle) // ture
```

## Storing subscribe data in a struct
```go
var subscriber struct {
	name string
	rate float64
	active bool
}
subscriber.name = "Aman Singh"
subscriber.rate = 4.99
subscriber.active = true
fmt.Println("Name:", subscriber.name) // Name: Aman Singh
fmt.Println("Monthly rate:", subscriber.rate) // Monthly rate: 4.99
fmt.Println("Active?", subscriber.active) // Active? true
```

## Defined types and structs
- **Type definitions** allow you to create types of your own.
- They let you create a new **defined type** that is based on an **underlying type**
- To write a type defintion: `type myType struct { ... }`
- Just like variables, type definitions can be written within a function..
  - That will limit its scope to that function's block.
  - Types are usually defined outside of any functions, at the package level.

```go
package main

import "fmt"

type part struct {
	description string
	count int
}

type car struct {
	name string
	topSpeed float64
}

func main() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name) // Name: Porsche 911 R
	fmt.Println("Top speed:", porsche.topSpeed) // Top speed: 323

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:", bolts.description) // Description: Hex bolts
	fmt.Println("Count:", bolts.count) // Count: 24
}
```

## Using a defined type for magazine subscribers
```go
package main

import "fmt"

type subscriber struct {
	name string
	rate flot64
	active bool
}

func main() {
	var subscriber1 subscriber
	subscriber1.name = "Aman Singh"
	fmt.Println("Name:", subscriber1.name) // Name: Aman Singh
	var subscriber2 subscriber
	subscriber2.name = "Beth Ryan"
	fmt.Println("Name:", subscriber2.name) // Name: Beth Ryan
}
```

## Using defined types with functions
```go
package main

import "fmt"

type part struct { 
	description string
	count int
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func minimumOrder(description string) part {
	var p part
	p.descriptioin = description
	p.count = 100
	return p
}

func main() {
	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts) // Description: Hex bolts
									// Count: 24
	p := minimumOrder("Hex bolts")
	fmt.Println(p.description, p.count) // Hex bolts 100
}
```

```go
package main

import "fmt"

type subscriber struct {
	name string
	rate float64
	active bool
}

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rage:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	subscriber1.rate = 4.99
	printInfo(subscriber1)	// Name: Aman Singh
													// Monthly rate: 4.99
													// Active? true
	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)	// Name: Beth Ryan
													// Monthly rate: 5.99
													// Active? true
}
```

## Modifying a struct using a function
- Go is a "pass-by-value" language
  - Function parameters receive a copy of the arguments the function was called with.
  - If a function changes a parameter value, it is changing the copy, not the original.
- The solution is to have the function parameter accept a pointer to a value, instead of accepting a value directly.

```go
package main

import "fmt"

type subscriber struct {
	name string
	rate float64
	active bool
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)	// 4.99
}
```

## Accessing struct fields through a pointer
- Use the `*` operator ("value-at" operator) to get the value at the pointer.
- Just putting a `*` before the struct pointer won't work.
  - If you write `*pointer.myField`, Go thinks that `myField` must contain a pointer. => It doesn't => error.
  - To get this to work, you need wrap `*pointer` in parantheses.
  - That will case the `myStruct` value to be retrieved, after which you can access the struct field.

```go
func main {
	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	fmt.Println((*pointer).myField)	// 3
}
```

- The dot operator lets you access field via pointers to structs, just as you can access fields directly from struct values.
  - You can leave off the parentheses and the `*` operator.

```go
func main() {
	var value myStruct
	value.myField = 3
	var pointer *myStruct = &value
	fmt.Println(pointer.myField)	// 3
}
```

- Works for assigning to struct fields through a pointer as well.

```go
func main() {
	var value myStruct
	var pointer *myStruct = &value
	pointer.myField = 9
	fmt.Println(pointer.myField)	// 9
}
```

## Pass large structs using pointers
- functions receive a copy of the arguments they are called with, even if they are a big value like a struct.
- That is why, unless your struct has only a couple small fields, it is often a good idea to pass functions a *pointer* to a struct, rather than the struct itself.
  - This is true even if the function doesn't need to modify the struct.
  - W

```go
// code above here omitted
type subscriber struct {
	name string
	rate float64
	active bool
}

func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	applyDiscount(subscriber1)
	printInfo(subscriber1)	// Name: Aman Singh
													// Monthly rate: 4.99
													// Active? true
	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)	// Name: Beth Ryan
													// Monthly rate: 5.99
													// Active? true
}
```

## Moving our struct type to a different package

## A defined type's name must be capitalized to be exported

## Struct field names must also be capitalized to be exported

## Struct literals
```go
myCar := car{name: "Corvette", topSpeed: 337}
```

- You can omit some or even all of the fields from the curly braces.
  - Omited fields will be set to the zero value for their type.

## Creating an Employee struct type

## Creating an Address struct type

## Adding a struct as a field on another type

## Anonymous struct fields
- Go allows you to define anonymous fields: struct fields that have no name of their own, just a type.
  - We can use an anonymous field ot make our inner struct easier to access.

```go
package magazine

type Subscriber struct {
	Name string
	Rate float64
	Active bool
	Address // Anonymous
}

type Employee struct {
	Name string
	Salary float64
	Address // Anonymous
}

type Address struct {
	// Fields omitted
}
```

- When you declare an anonymous field, you can use the field's tyep as if it were the name of the field.

## Embedding structs
- An inner struct that is stored within an outer struct using an anonymous field is said to be **embedded** within the outer struct.
  - Fields for an embedded struct are **promoted** to the outer struct, meaning you can access them as if they belong to the outer struct.
- `Address` struct type is embedded within the `Subscriber` and `Employee` struct types.
  - You don't have to write `subscriber.Address.City`
  - You can just write `subscriber.City`