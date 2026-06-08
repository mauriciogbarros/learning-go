# Chapter 4 - Packages
## Different programs, same function
- If parts of your code are shared between multiple programs, you should consider moving them into packages.

## The Go workspace directory holds package code
- Go tools look for package code in a special directory called **workspace**.
  - By default, the workspace is a directory named *go* in the current user's home directory.
- Contains three subdirectories:
  - *bin*, which holds compiled binary executable programs.
  - *pkg*, which holds compiled binary package files.
  - *src*, which holds Go source code.

## Creating a new package
- Within the *go* directory, create a directory named *src*.
- By convention, the package's directory, should have the same name as a package.

## Pool Puzzle
```go
/*
(user's home directory)
└── go
		├── bin
		├── pkg
		└── src
				└── calc
						└── calc.go
*/

package calc

func Add(first float64, second float64) float64 {
	return first + second
}

func Subtract(first float64, second float64) float64 {
	return first - second
}

package main

import (
	"calc"
	"fmt"
)

func main() {
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Subtract(7, 3))
}
```

## Package naming conventions
- A package name should be all lowercase.
- The name should be abbreviated if the meaning is fairly obvious (such as `fmt`).
- It should be one word, if possible. If two words are needed, they should not be separated by underscores, and the second word should *not* be capitalized. (The `strconv` package is one example.)
- Imported package names can conflict with local variable names, so don't use a name that package users are likely to want to use as well.
  - For example, if the `fmt` package were named `format`, anyone who imported that package would risk conflicts if they named a local variable `format`.

## Package qualifiers
- When accessing a function, variable, or the like that's exported from a different package, you need to qualify the name of the function or variable by typing the package name before it.
- When you access a function or variable that's defined in the *current* package, however, you should *not* qualify the package name.

## Constants
- Many packages export **constants**: named values that never change.
  - `const` keyword
  - A value must be assigned during declaration
  - The operator `:=` is not applicable
  - As with variable declarations, the type can be omitted, and it will be inferred from the value being assigned.
  - As with variables and functions, a constant whose name begins with a capital letter is exported, and we can access it from other packages by qualifying its name.

```go
const TriangleSides int = 3
```

## Nested package directories and import paths
- Some sets of packages are grouped together by import path prefixes like "archive/" and "math/".
  - These import path prefixes *are* created using directories.
- When you import a package, only the source code file stored directly under that directory will be imported, not the subdirectories.

## Installing program executables with "go install"
- When we use `go run`
  - o has to compile the program as well as the packages it depends on before it can execute it.
  - It throws that compiled code away when it's done.
- `go build` saves an executable binary file in the current directory.
  - Risk of littering your Go workspace with executables in random, inconvenient places.
- The `go install` command also saves compiled binary versions.
  - But in the `bin` directory of your Go workspace.

## Changing workspaces with the GOPATH environment variable
- By default, the Go workspace is a directory named *go* in the current user's home directory.

## Setting GOPATH
- You can change the workspace by setting the `GOPATH` environment variable to the path of a different directory.

## Publishing packages
- In GitHub, break up the URL and use the pieces as directory names.

```go
package main

import (
	"fmt"
	"github.com/username/repository/keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nill {
		lot.Fatal(err)
	}

	// More code here
}
```

## Downloading and installing packages with "go get"
- To download and install a package, you can use the `go get` command.

```bash
go get github.com/headfirstgo/greeting
```

- The go tool will connect to github.com, download the Git repository at the */headfirstgo/greeting* path, and save it in your Go workspace *src* directory.
- The `go get` command will automatically create whatever subdirectories are needed to set up the appropriate import path.
- The packages are then ready to be imported.

```go
import (
	"github.com/headfirstgo/greeting"
	"github.com/headfirstgo/greeting/dansk"
	"github.com/headfirstgo/greeting/deutsch"
)
```

## Reading package documentation with "go doc"
- You can use the `go doc` command to display documentation on any package or function.

## Documenting your packages with doc comments
- Ordinary Go comments that appear immediately before a package clause or function declaration are treated as doc comments, and will be displayed in `go doc` output.
- There are few conventions to follow when adding doc comments:
  - Comments should be complete sentences.
  - Package comments should begin with "Package" followed by the package name: `// Package mypackage enables widget management.`
  - Function comments should begin with the name of the function they describe: `// MyFunction converts widgets to gizmos.`
  - Other than indentation for code samples, don't add extra punctuation characters for emphasis or formatting. Doc comments will be displayed as plain text, and should be formatted that way.

## Serving HTML documentation to yourself with "godoc"
- The same software that powers the *golang.org* site's documentation section is actually available on your computer too.
  - It's a tool called `godoc`, and it's automatically installed along with Go.
  - The `godoc` tool generates HTML documentation based on the code ini your main Go installation and your workspace.
  - It includes a web server that can share the resulting pages with browsers.