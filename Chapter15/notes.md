# Chapter 15. Responding to requests: Web Apps
## Writing web apps in Go
- `net/http` package

## Browsers, requesets, servers, and responses
- Today it is much more common for a server to communicate with a program to fulfill the requests, instead of reading from a file.

## A simple web app
- `net/http`
  - Make requests to servers
  - Respond to requests
- Example: `hello_web`
  - Run the server: `go run hello_web.go`
  - Open browser and enter the following URL: `http://localhost:8080/hello`
  - The browser will send a request to the app, which will respond with `"Hello, Web!"`.

## Your computer is talking to itself
- `localhost`: tell the browser that it needs to establish a connection from the computer to the same computer.
- `8080`: port, a numbered network communication channel that an application can listen for messages on.

## Our simple web app, explained
- `http.HandleFunc` with the string `"/hello"` and the `viewHandler` function
  - Tells the app to call `viewHandler` whenever a request for a URL ending in `/hello` is received.
- `http.ListenAndServe`: starts up the web server.
  - `localhost:8080` causes the server to accept requests only from your own machine on port 8080.
  - `nil` value in the second argument just means that requests will be handled using functions set up via `HandleFunc`
  - It is called **after** `HandleFunc` because it will run forever, unless it encounters an error.
- `viewHandler`:
  - the server passes `viewHandler`:
    - `http.ResponseWriter`, which is used for writing data to the browser response
    - a pointer to `http.Request` value, which represents the browser's request.
  - Within `viewHandler`, we add data to the response by calling the `write` method on the `ResponseWriter`.
    - `Write` doesn't accept strings, so we convert `"Hello, web!"` string to a `[]byte`
    - `Write` returns the number of bytes successfully written, and any error encountered.

## Resource paths
- The part of a URL following the host address and port is the resource path.
  - It tells the server which of its many resources you want to act on.

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
|x|http://localhost:4567/c|
|y|http://localhost:4567/a|
|z|http://localhost:4567/b|


## First-class functions
- When we call `http.HandleFunc` with handler functions, we are **not** calling the handler function and passing its result to `HandleFunc`.
- We are passing the **function itself** to `HandleFunc`.
  - That function is stored to be called later when a matching request path is received.
- In a programming language with first-class functions, functions can be assigned to variables, and then called from those variables.

```go
func sayHi() {
	fmt.Println("Hi")
}

func main() {
	var myFunction func() // Declare a variable with a type of func(),
												// that can hold a function.
	myFunction = sayHi		// Assign the sayHi function to the variable.
	myFunction()					// Call the function stored in the variable.
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

```go
func sayHi() {
	fmt.Println("Hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	var greeterFunction func()
	var mathFunction func(int, int) float64
	greeterFunction = sayHi
	mathFunction = divide
	greeterFunction()
	fmt.Println(mathFunction(5, 2))
}
```

- Functions that accept a function as parameter also need to specify the parameters and return types the passed-in function should have.