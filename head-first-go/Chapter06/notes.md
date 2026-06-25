# Chapter 6. appending issue: Slices
- Slices are a collection type that can grow to hold additional items.

## Slices
- Like arrays, slices are made up of multiple elements, all of the same type.
- Unlike arrays, functions are available that allow us to add extra elements onto the end of a slice.
- To declare: `var mySlice []string`

- Unlike with array variables, declaring a slice variable doesn't automatically create a slice.
  - Fhat that, you can call the built-in `make` function.

```go
var notes []string
notes = make([]string, 7) // create a slice with seven strings
```

- You don't have to declare the variable and create the slice in separate steps

```go
primes := make([]int, 5)
```

- The built-in `len` function works the same way with slices as it does with arrays.

```go
notes := make([]string, 7)
primes := make([]int, 5)
fmt.Println(len(notes)) // 7
fmt.Println(len(primes)) // 5
```

- Both `for` and `for ... range` loops work just the same

```go
letters := []string{"a", "b", "c"}
for i := 0; i < len(letters); i++ {
	fmt.Println(letters[i])
}

for _, letter := range letters {
	fmt.Println(letter)
}
```

## Slice literals
- Just like with arrays, if you know in advance values a slice will start with, you can initialize the slice with those values using a **slice literal**.
  - There is no need to call the `make` function; usisng a slice literal in your code will create the slice *and* prepopulate it.
  - `[]int{9, 18, 27}`

>Slieces are built on top of arrays.

## The slice operator
- Every slice is built on top of an underlying array.
  - It's the underlying array that actually holds the slice's data
  - The slice is merely a view into some (or all) of the array's elements.
- When you use the `make` function or a slice literal to create a slice, the underlying array is created for you automatically.
- You can also create the array yourself, and then create a slice based on it with the **slice operator**.
  - `mySlice := myArray[1:3]`
  - 1: Index of array where slice should start
  - 3: Index of array slice should stop before

```go
underlyingArray := [5]string{"a", "b", "c", "d", "e"}
slice1 := underlyingArray[0:3]
fmt.Println(slice1) // [a b c]
```

- To include the last element of an underlying array, specify a second index that is one *beyond* the end of the array.

```go
slice2 := underylyingArray[2:5]
fmt.Println(slice2) // [c d e]
```

- Make sure you don't go any further than that, though, or you will get an error.

- The slice operator has defaults for both the start and stop indexes.
  - If you omit the start index, a value of `0` (the first element of the array) will be used.
  - If you omit the second index, everything from the start index to the end of the underlying array will be included in the resulting slice.

## Underlying arrays

## Change the underlying array, change the slice
- With `make` and with slice literals, you never have to work with the underlying array.

## Add onto a slice with the "append" function
- Go offers a built-in `append` function that takes a slice, and one or more values you want to append to the end of that slice.
  - It returns a new, larger slice with all the same elements as the original slice, plus the new elements added onto the end.

```go
slice := []string{"a", "b"}
fmt.Println(slice, len(slice)) // [a b] 2
slice = append(slice, "c")
fmt.Println(slice, len(slice)) // [a b c] 3
slice = append(slice, "d", "e")
fmt.Println(slice, len(slice)) // [a b c d e] 5
```

- A slice's underlying array can't grow in size.
  - If there isn't room in the array to add elements, all its elements will be copied to a new, larger array, and the slice will be updated to refer to this new array.
  - All this happens behind the scenes in the `append` function.
  - There no easy way to tell whether the slice returned from `append` has the *same* underlying array as the slice passed in, or a *different* underlying array.
  - If you keep both slices, this can lead to some unpredictable behavior.

## Slices and zero values
- As with arrays, if you access a slice element that no value has bee assigned to, you will get the zero value for that type back.
- Unlike arrays, the slice variable itself also has a zero value.
  - It's `nil`.
  - That is, a slice variable that no slice has been assigned to will have a value of `nil`.
- The `len` function will return `0` if it's passed a `nil` slice.
- The `append` function also treats `nil` slices like empty slices.
  - If you pass an empty slice to `append`, it will add the item you specify to the slice, and return a slice with one item.
  - If you pass a `nil` slice to `append`, you will also get a slice with one item back.

## Reading additional file lines using slices and "append"

## Trying our improved program

## Returning a nil slice in the event of an error

## Command-line arguments

## Getting command-line arguments from the os.Args slice
- The name of the executable is included as the first element of `os.Args`
- If we use a slice operator of `[1:]` on `os.Args`, it will give us a new slice that omits the first element (whose index is `0`) and includes the second element (index `1`) through the end of the slice.

## Updating our program to use command-line arguments

## Variadic functions
- `Println` and `append` are declared as variadic functions, for example.
- A **variadic function** is one that can be called with a *varying* number of arguments.
- To make a function variadic, use an ellipsis (`...`) before the type of the last (or only) function parameter in the function declaration.

```go
func myFunc(param1 int, param2 ...string) {
	// function code here
}
```

- The last parameter of a variadic function receives the variadic arguments as a slice, which the function can then process like any other slice.

- A function can toke one or more nonvariadic arguments as well.
  - Although a function caller can omit variadic arguments (resulting in an empty slice), nonvariadic arguments are always required.
  - It's a compile error to omit those.
  - Only the *last* parameter in a function definition can be variadic.
  - You can't place it in front of required parameters.

## Using variadic functions

## Code Magnets

## Using a variadic function to calculate averages

## Passing slices to variadic functions
- Go provides special syntax for this situation.
- When calling a variadic function, simply add an ellipsis (`...`) following the slice you want to use in place of variadic arguments.