# Chapter 5. on the list: Arrays
## Arrays hold collections of values
- An **array** is a collection of values that all share the same type.
  - The values an array holds are called its **elements**.
  - An array holds a specific number of element, and it cannot grow or shrink.
  - To declare a variable that holds an array: `var myArray [4]string`

- Example:
```go
var primes[5]int
primes[0] = 2
primes[1] = 3
fmt.Println(primes[0]) // 2
```

## Zero values in arrays
- As with variables, when an array is created, all the values it contians are initialized to the zero values for the type that array holds.

## Array literals
```go
var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
fmt.Println(notes[3], notes[6], notes[0]) // fa ti do
primes := [5]int{2, 3, 5, 7, 11}
fmt.Println(primes[0], primes[2], primes[4]) // 2 5 11
```

- You can spread array literals over multiple lines, but you are required to use a comma before each newline character in your code.
  - Including the final entry in the array literal.

## Functions in the "fmt" package know how to handle arrays
```go
var notes [3]string = [3]string{"do", "re", "mi"}
var primes [5]iint = [5]int{2, 3, 5, 7, 11}
fmt.Println(notes) // [do re mi]
fmt.Println(primes) // [2 3 5 7 11]
```

- Using `%#v` verb used by the `Printf` and `Sprintf` functions, which formats values as they would appear in Go code.

```go
fmt.Printf("%#v\n", notes) // [3]string{"do", "re", "mi"}
fmt.Printf("%#v\n", primes) // [5]int{2, 3, 5, 7, 11}
```

## Accessing array elements within a loop
```go
notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
for i := 0; i <= 2; i++ {
	fmt.Println(i, notes[i])
}
// 0 do
// 1 re
// 2 mi
```

- Trying to access an index that is outside the array will cause a **panic**.
  - An error that occurs while your program is running.
  - Normally, a panic causes your program to crash and display an error message to the user.

## Checking array length with the "len" function
```go
primes := [5]int{2, 3, 5, 7, 11}
fmt.Println(len(primes)) // 5
```

## Looping over arrays safely with "for ... range"
```go
for index, value := range myArray {
	// Loop block here
}
```

## Using the blank identifier with "for ... range" loops
```go
for _, note := range notes {
	fmt.Println(note)
}
```

## Reading a text file
