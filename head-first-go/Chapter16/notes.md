# Chapter 16. A pattern to follow: HTML templates
- Go offers the `html/template` package, a powerful way to include data in your app's HTML responses.

## A guestbook app
- Respond to requests for the main guestbook page.
- Format the response using HTML.
- Fill the HTML page with signatures.
- Set up a form for adding a new signature.
- Save submitted signatures.

## Functions to handle a request and check errors

## Making a signature list in HTML

## Making our app respond with HTML

## The "text/template" package
- The `html/tempalte` package is based on the `text/template` package.
  - You work with the two packages in almost exactly the same way, but the `html/template` has some extra security features needed for working with HTML.
- The program below uses `text/template` to parse and print a template string.

```go
package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	text := "Here's my template!\n"

	// Create a new template value based on the text
	tmpl, err := template.New("test").Parse(text)
	check(err)

	// Write out the template text.
	err = tmpl.Execute(os.Stdout, nil)
	check(err)
}
```

## ResponseWriters and os.Stdout both satisfy io.Writer

## Inserting data into templates using actions
- The second parameter to a `Template` value's `Execute` method allows you to pass in data to insert in the template.
- To insert data in a template, you add **actions*8 to the template text.
  - Actions are denoted with double curly braces, `{{}}`.
  - Inside the double braces, you specify data you want to insert or an operations you want the template to perform.
  - Whenever the template encounters and action, it will evaluate its contents, and insert the result into the template text in place of the action.

```go
func main() {
	templateText := "Template start\nAction: {{.}}\nTemplate end\n"
	tmpl, err := template.New("test").Parse(templateText)
	check(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)
	err = tmpl.Execute(os.Stdout, 42)
	check(err)
	err = tmpl.Execute(os.Stdout, true)
}
```

- Let's set up an `executeTemplate` function that will let us experiment with them more easily.

```go
func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n", 123.5)
}
```

## Making parts of a template optional with "if" actions
- A section of a template between an `{{if}}` action and its corresponding `{{end}}` marker will be included only if a condition is true.

```go
executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", true)
executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", false)
```

## Repeating parts of a template with "range" actions
- A section of a template between `{{range}}` action and its corresponding `{{end}}` marker will be repeated for each value collected in an array, slice, map, or channel.
- Within the repeated section, the value of dot will be set to the current element from the collection

```go
templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
executeTemplate(templateText, []string{"do", "re", "mi"})
```

## Inserting struct fields into a template with actions
- Simple types usually can't hold the variety of information needed to fill in a template.
  - It's more common to use struct types when executing a template.
- If the value in dot is a struct, then an action with dot followed by a field name will insert that field's value in the template.

```go
type Part struct {
	Name string
	Count int
}

templateText := "Name: {{.Name}}\nCount: {{.Count}}\n"
executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
executeTemplate(templateText, Part{Name: "Cables", Count 2})
```

```go
type Subscriber struct {
	Name string
	Rate float64
	Active bool
}

templateText = "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
subscriber := Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
executeTemplate(templateText, subscriber)
subscriber = Subscriber(Name: "Joy Carr", Rate: 5.99, Active: false)
executeTemplate(templateText, subscriber)
```

## Reading a slice of signatures in from a file
```go
import (
	"buffio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Code omitted...

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

// Code omitted

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getString("signatures.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFile("view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}
```

## A struct to hold the signatures and signature count
```go
type Guestbook struct {
	SignatureCount int
	Signatures []string
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures: signatures,
	}
	err = html.Execute(writer, guestbook)
	check(err)
}
```

## Updating our template to include our signatures
```html
<h1>Guestbook</h1>
<div>
	{{.SignatureCount}} total signatures - 
	<a href="/guestbook/new">Add Your Signature</a>
</div>

<div>
	{{range .Signatures}}
		<p>{{.}}</p>
	{{end}}
</div>
```

## Letting users add data with HTML forms

## Responding with the HTML form
- Add a `newHandler` function
  - It will look much like the early versions of our `viewHandler` function.
  - `newHandler` should take an `http.ResponseWriter` and a pointer to an `http.Request` as parameters.
  - It should call `template.ParseFiles` on the **new.html** file.
  - And then it should call `Execute` on the resulting template, so that the contents of **new.html** get written to the HTTP response.
  - We won't be inserting any data into this template, so we pass `nil` as the data value for the call to `Execute`.
  - Then we need to ensure that the `newHandler` function is called when the "Add Your Signature" link is clicked.

## Form submission requests
- The browser will just make another request for the **/guestbook/new** path.

## Path and HTTP method for form submissions
- Submitting a form actually requires two requests to the server
  - One to get the form
  - A second to send the user's entries back to the server.

## Getting values of form fields from the request
- The form data is embedded in the request itself, rather than being appended to the request path as a parameter.
- Add a request handler function for the path of `/guestbook/create` to a new `createHanlder` function.
  - It should accept an `http.ResponseWriter` and a pointer to an `http.Request`, just like the other handler functions.
  - Unlike the other handler functions, data can be accessed through the `http.Request` pointer that gets passed.
  - Call the `FormValue` method on the `http.Request`, and pass it the string `"signature"`, which will return a string with the value of the `"signature"` form field, and store it in a variable named `signature`.

## Saving the form data
- `os.OpenFile`
  - Opens the **signatures.txt** file, creating it if it doesn't exist.
  - Adds a line of text to the end of the file.
  - Closes the file

## HTTP redirects
- `http.Redirect`, which sends a response to the browser directing it to load a different resource than the one it requested.
- It takes:
  - `http.ResponseWriter` and `*http.Request`
  - A string with a path to redirect the browser to; in this case `"/guestbook"`
  - A status code to give the browser, every http response needs to include a status code. We'll use `http.StatusFound`, which will cause the redirect response to have a status of 302 ("Found").