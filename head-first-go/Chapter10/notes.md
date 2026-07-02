# Chapter 10. keep it to yourself: Encapsulation and Embedding
- **Encapsulation**: a way to protect your struct type's fields from that invalid data.
- **Embed** other types within your struct type.

## Creating a Data struct type
- Defined types can use any other type as their underlying type, including structs.
- Creating a `Date` struct typ to hold year, month, and day values.
  - `Year`, `Month`, and `Day` are fields of the struct, each with a type of `int`.

```go
package main

import "fmt"

type Date struct {
	Year int
	Month int
	Day int
}

func main() {
	date := Date{Year: 2019, Month: 5, Day: 27}
	fmt.Println(date) // {2019 5 27}
}
```

## People are setting the Date struct field to invalid values!
- Data validation

## Setter methods
- By convention, Go setter methods are usually named in the form `SetX`, where `X` is the thing that is being set.

## Setter methods need pointer receivers

## Adding the remaining setter methods

## Adding validation to the setter methods

## The fields can still be set to invalid values!

## Moving the Date type to another package

## Making Date fiels unexported

## Accessing unexported fields through exported methods
- Unexported variables, struct fields, functions, methods, and the like can still be accessed by exported functions and methods in the same package.

## Getter methods

## Encapsulation
- The practice of hiding data in one part of a program from code in another part is known as **encapsulation**.
- In Go, data is encapsulated within packages, using unexported variables, struct fields, functions, or methods.
- Go developers generally only rely on encapsulation when it is necessary, such as when field data needs to be validated by setter methods.

## Embedding the Date type in an Event type
- Create another file within the `calendar` package folder, named `event.go`.
- Within that file, define an `Event` type with two fields:
  - `Title` with a type of `string`
  - an anonymous `Date` field.

## Unexported fields don't get promoted
- Embedding `Date` in the `Event` type will not cause the `Date` fields to be promoted to the `Event`, though.
- The `Date` fields are unexported, and Go does not promote unexported fields to the enclosing type.

## Exported methods get promoted just like fields
- As with unexported fields, unexported methods are not promoted.

## Encapsulating the Event Title field

## Promoted methods live alongside the outer type's methods
