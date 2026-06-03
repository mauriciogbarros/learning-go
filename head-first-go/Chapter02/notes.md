# Chapter 2 - Conditionals and loops
## Calling methods
[time example](./example_time/example_time.go)
[string fixer example](./example_stringFixer/example_stringFixer.go)

## Comments
`//` => single line
`/* */` => multi-line

## Multiple return values from a function or method
- In Go, functions can return any number of values.
- The most common use of multiple return values in Go is to return an additional error value that can be consulted to find out if anything went wrong while the function or method was running.

```go
bool, err := strconv.ParseBool("true")
file, err := os.Open("myfile.txt")
response, err := http.Get("http://golang.org")
```
- However, Go doesn't allow us to declare a variable unless we use it.

### Option 1: Ignore the error return value with the blank identifier
- `_`: **blank identifier**
  - Essentially discards it

```go
input, _ := reader.ReadString('\n')
```

### Option 2: Handle the error
- The `log` package has a `Fatal` function that can log a message to the terminal and stop the program.

## Conditionals
- No parenthesis around conditionals

```go
if grade == 100 {
	fmt.Println("Perfect!")
} else if grade >= 60 {
	fmt.Println("You pass.")
} else {
	fmt.Println("You fail!")
}
```

## Avoid shadowing names

## Converting strings to numbers

## Only one variable in a short variable declaration has to be new