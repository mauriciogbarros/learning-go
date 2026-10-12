# Chapter 15. responding to requests: Web Apps
## Writing web apps in Go
- With `net/http` package

## Browsers, requests, servers, and responses
- Today, it is much more common for the server to communicate with a program to fulfill the requests, instead of reading from a file.

## A simple web app
- The `net/http` package also includes a small web server, so it is also able to **respond** to requests.

```go
package main

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	// writer: a value for updating the response that will be sent to the
	//				 browser.
	// request: a value representing the request from the browser.

	message := []byte("Hello, web!")
	_, err := writer.Write(message) // Add "Hello, web!" to the response
		// The Write method returns the number of bytes successfully
		// written, and any error encountered.

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// If we receive a request for a URL ending in "/hello"...
	// ... then call the viewHandler function to generate a response.
	http.HandleFunc("/hello", viewHandler)

	// Listen for browser requests and response to them.
	err != http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
```

## Your computer is talking to itself
- localhost

## Our simple web app, explained

## Resource paths
- The `net/http` server pulls the path off the end of the URL, and uses it in handling the request.

## Responding differently for different resource paths
- Update the example to provide greetings in three different languages.

```go
package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello, web!")
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Salut web!")
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Namaste, web!")
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/salut", frenchHandler)
	http.HandleFunc("namaste", hindiHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
```

## Exercise

```go
package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func d(writer http.ResponseWriter, request *http.Request) {
	write(writer, "z")
}

func e(writer http.ResponseWriter, request *http.Request) {
	write(writer, "x")
}

func f(writer http.ResponseWriter, request *http.Request) {
	write(writer, "y")
}

func main() {
	http.HandleFunc("/a", f)
	http.HandleFunc("/b", d)
	http.HandleFunc("/c", e)
	err := http.ListenAndServe("localhost:4567", nil)
	log.Fatal(err)
}
```

|Response|URL to generate response|
|:---:|:---|
|x|localhost:4567/c|
|y|localhost:4567/a|
|z|localhost:4567/b|

## First-class functions
- The Go language supports **first-class functions**
- In a programming language with first-class functions, functions can be assigned to variables, and then called from those variables.

```go
func sayHit() {
	fmt.Println("Hi")
}

func main() {
	// Declare a variable with a type of "func()".
	// This variable can hold a function.
	var myFunction func()

	// Assign the sayHi function to the variable
	myFunction = sayHi

	// Call the function stored in the variable
	myFunction() // Hi
}
```

## Passing functions to other functions
- Programming languages with first-class functions also allow you to pass functions as arguments to other functions.

```go
func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func main() {
	twice(sayHi)
	twice(sayBye)
}
```

## Functions as types
- A function's parameters and return value are part of its type.
- A variable that holds a function needs to specify what parameters and return values that function should have.
  - That variable can only hold functions whose number and types of parameters and return values match the specified type.

```go
func sayHi() {
	fmt.Println("Hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	// This variable will hold a function with no parameters and no return
	// value.
	var greeterFunction func()

	// This variable will hold a function with two int parameters and a
	// float64 return value.
	var mathFunction func(int, int) float64

	greeterFunction = sayHi
	mathFunction = divide

	greeterFunction()									// Hi
	fmt.Println(mathFunction(5, 2))		// 2.5
}
```

## Pool puzzle
functionC

functionB
bool
```go
func callFunction(passedFunction func()) {
	passedFunction()
}

func callTwice(passedFunction func()) {
	passedFunction()
	passedFunction()
}

func callWithArguments(passedFunction func(string, bool)) {
	passedFunction("This sentence is", false)
}

func printReturnValue(passedFunction func() string) {
	fmt.Println(passedFunction())
}

func functionA() {
	fmt.Println("function called")
}

func functionB() string {
	fmt.Println("function called")
	return "Returning from function"
}

func functionC(a string, b bool) {
	fmt.Println("function called")
	fmt.Println(a, b)
}

func main() {
	callFunction(functionA)
	callTwice(functionA)
	callWithArguments(functionC)
	printReturnValue(functionB)
}
```
