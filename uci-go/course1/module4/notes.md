# 4 - Protocols and Formats
## 1 - Overview
Packages to communicate with other systems

## 2 - RFCs
- **Requests for Comments** (RFC)
- Definitions of internet protocols and formats
- Example protocols
  - HTML (Hypertext Markup Language): 1866
  - URI (Uniform Resource Identifier): 3986
  - HTTP (Hypertext Transfer Protocols): 2616
- Golang provides packages for important RFCs
- Functions which encode and decode protocol format
- Example: "net/http"
  - Web communication protocol
  - `http.Get(www.uci.edu)`
- Example: "net"
  - TCP/IP and socket programming
  - `net.Dial("tcp", "uci.edu:80")`

### JSON
- JavaScript Object Notation
- RFC 7159
- Format to represent structured information
- **Attribute-value** pairs
  - struct or map
- Basic value types:
  - bool, number, string, array, "object"

### JSON Example
- Go struct

```go
p1 := Person(name: "Joe", addr: "a st.", phone: "123")
```

- Equivalent JSON object

```json
{
	"name": "joe",
	"addr": "a st.",
	"phone": "123"
}
```

## 3 - JSON
### JSON Properties
- All Unicode
- Human-readable
- Fairly compact representation
- Types can be combined recursively
  - Array of struct, struct in struct, ...

### JSON Marshalling
- Generating JSON representation from an object

```go
type struct Person {
	name string
	addr string
	phone string
}
```

- `json.Marshal()` returns JSON representation as []byte

```go
p1 := Person(name: "joe", addr: "a st.", phone: "123")
barr, err := json.Marshal(p1)
```

### JSON Unmarshalling
- `json.Unmarshal()` converts a JSON []byte into a Go object
  - Pointer to Go object is passed to `Unmarshal()`
  - Object must "fit" JSON []byte

```go
var p2 Person
err := json.Unmarshal(barr, &p2)
```

## 4 - File Access, ioutil (depecrated)
### File
- Linear access, not random access
  - Mechanical delay
- Basic operations
  - Open - get handle for access
  - Read - read bytes into []byte
  - Write - write []byte into file
  - Close - release handle
  - Seek - move read/write head

### ioutil File Read
- "io/ioutil" package has basic functions

```go
dat, e := ioutil.ReadFile("test.txt")
```
- `dat` is []byte filled with contents of entire file
- Explicit open/close are needed
- Large files case a problem

### ioutil File Write
```go
dat = "Hello, world"

err := ioutil.WriteFile("outfile.txt", dat, 0777)
```
- Writes []byte to file
- Creates a file
- Unix-style permission bytes

## 5 - File Access, os
### os Package File Access
- `os.Open()` opens a file
  - Returns a file descriptor (File)
- `os.Close()` closes a file
- `osRead()` reads from a file into a []byte
  - Fills the []byte
  - Control the amount read
- `os.Write()` writes a []byte into a file

### os File Reading
- Opening and reading

```go
f, err := os.Open("dt.txt")
barr := make([]byte, 10)
nb, err := f.Read(barr)
f.Close()
```
- Reads and fills `barr`
- `Read` returns # of bytes read
- May be less than []byte length

### os File Create/Write
```go
f, err := os.Create("outfile.txt")

barr := []byte{1, 2, 3}
nb, err := f.Write(barr)
nb, err := f.WriteString("Hi")
```
- `WriteString()` writes a string
- `Write()` writes a []byte
  - Any Unicode sequence