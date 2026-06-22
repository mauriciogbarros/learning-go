# 3 - Composite Data Types
## 1 - Arrays
- Fixed-length series of elements of a chosen type
- Elements accessed using subscript notation, `[ ]`
- Indices start at `0`.
- Elements initialized to zero value

```go
var x [5]int
x[0] = 2
fmt.Printf(x[1])
```

### Array Literal
- An array pre-defined with values
  - `var x [5]int = [5]{1, 2, 3, 4, 5}`
- Length of literal must be length of array
- `...` for size in array literal infers size from numberr of initializers
  - x := [...]int{1, 2, 3, 4}

### Iterating through arrays
- Use a for loop with the range keyword

```go
x := [3]int {1, 2, 3}

for i, v range x {
	fmt.Printff("ind %d, val %d", i, v)
}
```
- Range returns two values
  - Index and element at index

## 2 - Slices
- A "window" on an **underlying array*
- Variable size, up to the whole array
- **Pointer** indicates the start of the slice
- **Length** is the number of elements in the slice
- **Capacity** is maximum number of elements
  - From start of slice to end of array

### Slice Examples
```go
arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
s1 := arr[1:3]
s2 := arr[2:5]
```

### Length and Capacity
- `len()` function returns the length
- `cap()` function returns the capacity

```go
a1 := [3]string("a", "b", "c")
sli1 := a1[0:1]
fmt.Printf(len(sli1), cap(sli1))
```

### Accessing Slices
- Writing to a slice changes the underlying array
- Overlapping slices refer to the same array elements

### Slice Literals
- Can be used to initialize a slice
- Creates the underlying array and creates a slice to reference it
- Slice points to the start of the array, length is capacity

```go
sli := []int{1, 2, 3}
```

## 3 - Variable Slices
### Make
- Create a slice (and array) using `make()`
- 2-argument version: specify type and length/capacity
- Init. to zero, length = capacity
  - `sli = make([]int, 10)`
- 3-argument version: specify length and capacity separately
  - `sli = make([]int, 10, 15)`

### Append
- Size of a slice can be increased by `append()`
- Adds elements to the end of a slice
- Inserts into underlying array
- Increases size of array if necessary
  - `sli = make([]int, 0, 3)`
- Length of sli is 0
  - `sli = append(sli, 100)

### Hash Table
- Contains key/value pairs
  - SSN and email
  - GPS coordinates and address
- Each value is associated with a unique key
- **Hash function** is used to compute the slot for a key

### Tradeoffs of Hash Tables
- Advantages
  - Faster lookup than lists: constant-time vs linear-time
  - Arbitrary keys: not ints, like slices or arrays
- Disadvantages:
  - May have collisions: two keys hash to same slot

## 4 - Maps
- Implementation of a hash
- Use `make()` to create a map

```go
var idMap map[string]int
ideMap = make(map[string]int)
```

- May define a map literal
  - `idMap := map[string]int { "joe": 123 }`

### Accessing maps
- Referencing value with [key]
  - `fmt.Println(idMap["joe"])`
- Returns zero if key is not present
- Adding a key/value pair
  - `idMap["jane"] = 456`
- Delete a key/value pair
  - `delete(idMap, "joe")`

### Map functions
- Two-value assignment tests for existence of the key
  - `id, p := idMap["joe"]`
  - `id` is value, `p` is presence of key
- `len()` returns number of values
  - `fmt.Println(len(idMap))`

### Iterating through a map
- Use a for loop with the range keyword
- Two-value assignment with range

```go
for key, val := range idMap {
	fmt.Println(key, val)
}
```

## 5 - Structs
- Aggregate data type
- Groups together other objects of arbitrary type
- Example: Person Struct
  - Name, Address, phone
  - Option 1: have 3 separate variables, programmer remembers that they are related.
  - Option 2: make a single struct which contains all 3 vars.

### Struct example
```go
type struct Person {
	name string
	addr string
	phone string
}

var p1 Person
```
- Each property is a field
- `p1` contains values for all fields

### Accessing struct fields
- Use dot notation

```go
p1.name = "joe"
x = p1.addr
```

### Initializing structs
- Can use `new()`
- Initializes fields to zero
  - `p1 := new(Person)`
- Can initialize using a struct literal
  - `p1 := Person(name: "joe", addr: "a st.", phone: "123")`