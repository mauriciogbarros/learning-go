# Module 1 - Getting Started with Go Programming
## Overview
- Packages
  - Groups related code units.
  - The primary organizational unit of Go code.
  - Supports collaboration and code sharing.

## Advantages of Go
1. Code runs fast
2. Garbage collection
3. Simpler objects
4. Concurrency is efficient

### Software Translation
- Machine language: CPU instructions represented in binary.
  - Runs directly on the processor
- Assembly language: CPU instructions with mnemonics
  - Easier to read (still hard)
  - Equivalent to machine language (very close)
- High-level languages
  - Commonly used languages (C, C++, Java, Python, Go, etc.)
  - Much easier to use => More abstractions

>All software must be translated into the machine language of processor.

### Compiled vs Interpreted
- Compilation: translate instructions once before running the code
  - C, C++, Go, Java (partially)
  - Translation occurs only once, saves time.
- Interpretation
  - Translate instructions wile code is executed
    - Python, Java (partially)
    - Translation occurs every execution
    - Requires an interpreter
  - Slower

### Efficiency vs. Ease-of-Use
- Compiled code is fast
- Interpreters make coding easier
  - Manage memory automatically
  - Infer variable types
- Go is a good compromise
  - It is compiled
  - Also has some features of intepreted
  - Garbage collection

### Garbage Collection
- Automatic memory management
  - Where should memory be allocated?
  - When can memory be deallocated?
- Manual memory management is hard
  - Deallocate too early, false memory access
  - Deallocate too late, wasted memory
    - Memory leak
- Go include garbage colleciton
  - Typically only done by interpreters
  - Slows downs execution a bit

## Objects
### Object-Oriented Programming
- Organize your code through **encapsulation**
  - Group together data and functions which are related.
  - User-defined type which is specific to an application.

### Object Example
- Implementing an application performing geomerty in 3D
  - Many functions will operate on points
  - Each point has data: x, y, z
  - Points also have **functions**:
    - DistToOrigin(), Quadrant(), ...
  - Point **class** defines data and functions
  - Point **objects** are instances of class
- Objects in Go
  - Go does not use the term class
  - Go uses **structs** with associated methods
  - Simplified implementation of classes
    - No inheritance, no constructors, no generics
  - Easier to code
    - If you like these features, this is a disadvantage

### Concurrency
- Performance limits
  - **Moore's Law** used to help performance
    - The number of transistors doubles every 18 months
  - More transistors used to lead to higher clock frequencies
  - **Power**/**temperature** constraints limit clock frequencies now
- Parallelism
  - Number of cores still increases over time
  - Multiple tasks may be performed at the same time on different cores
  - Difficulties with parallelism
    - When do tasks start/stop?
    - What if one task needs data from another task?
    - Do tasks conflict in memory?
- Concurrent Programming
  - **Concurrency** is the management of multiple tasks at the same time
  - Key requirement for large systems
  - Concurrent programming enables parallelism
    - Management of task execution
    - Communication between tasks
    - Synchronization between tasks
- Concurrency in Go
  - Go includes concurrency primitives
  - **Goroutines** represent concurrent tasks
  - **Channels** are used to communicate between tasks
  - **Select** enables task synchronization
  - Concurrency primitives are efficient and easy to use

### Workspaces & Packages
- Workspaces
  - Go files
  - Hierarchy of directories
    - Common organization is good for sharing
  - Recommended **three subdirectories**
    - src - source code files
    - pkg - packages (libraries)
    - bin - executables
  - Programmer typically has one workspace for many projects
- Workspace Information
  - Recommended not enforced
  - Workspace directory defined by **GOPATH** environment variable
    - Defined during installation
  - Go tools assums that code is in GOPATH
- Packages
  - Group of related source files
  - Each package can be imported by other packages
  - First line of file names the package
  - There must be one package called `main`
  - Building the main package generated an executable program
  - Main package needs a `main()` function
    - Where code execution starts

```go
package main
import "fmt"
func main() {
	fmt.Printf("hello, world\n")
}
```

### Go Tool
- Import
  - `import` keyword is used to access other packages
  - Go standard library includes many packages
    - "fmt"
  - Searches directories specified by GOROOT and GOPATH
- The Go Tool
  - A tool to manage Go source code
  - Several commands
  - `go build`: compiles the program
    - Arguments can be a list of packages or a list of .go files
    - Creates an executable for the main package, same name as the first .go file
    - .ee suffix for executable in Windows
  - `go doc`: prints documentation for a package
  - `go fmt`: formats source code files
  - `go get`: downloads packages an installs them
  - `go list`: list all installed packages
  - `go run`: compiles .go files and runs the executable
  - `go test`: runs tests using files ending in `_test.go`

### Variables
- Naming
  - Names are needed to refer to programming constructs
    - Variables, functions, ...
  - Names must start with a letter
  - Any number of letters, digits, underscores
  - Case sensitive
  - Don't use keywords
    - `if`, `case`, `package`, ...
- Variables
  - Data stored in memory
  - Must have a name and a type
  - All variables must have **declarations**
  - Most basic: `var x int`
  - Can declare many on the same line: `var x, y int`
- Variable Types
  - Type defines the values a variable may take and operations that can be performed on it.
  - Integer
    - Only integral values
    - Integer arithmetic
  - Floating point
    - Fractional (decimal) values
    - Floating point arithmetic
  - Strings
    - Byte (character) sequences
    - String comparison, search, concatenation, ...

### Variable Initialization
- Type declarations
  - Defining an alias (alternate name) for a type
  - May improve clarity: `type Celsius float64`, `type IDnum int`
  - Can delcare variables using the type alias: `var temp Celsius`, `var pid IDnum`
- Initializing variables
  - Initialize in the declaration
    - `var x int = 100`
    - `var x = 100` (inferrence)
  - Initialize after the declaration
    - `var x int`, `x = 100`
  - Uninitialized variables have a zero value
```go
		var x int // x = 0
		var x string // x = ""
```
- Short variable declaration
  - Can perform a declaration and initialization together with the **:=** operator
  - `x := 100`
  - Variable is declared as type of expression on the right hand side
  - Can only be done inside a function