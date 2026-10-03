# Module 3: Object orientation in Go
## 1 - Classes and Encapsulation
### Clases
- Collection of data fields and functions that share a well-defined responsibility
- Example: **Point** class
  - Used in a geometry program
  - Data: x coordinate, y coordinate
  - Functions:
    - `DistToOrigin()`
    - `Quadrant()`
    - `AddXOffSet()`
    - `AddYOffset()`
    - `SetX()`
    - `SetY()`
- Classes are a **template**
- Contain **data fields**, not data

### Object
- Instance of a class
- Contains real data

### Encapsulation
- Data can be protected from the programmer
- Data can be accessed only using methods
- Maybe we don't trust the programmer to keep data consistent

## 2 - Support for classes
### No "class" keyword
- Most OO languages have a class keyword
- Data fields and methods are defined inside a class block

```py
class Point:
	def __init__(self, xval, yval):
		 self.x = xval
		 self.y = yval
```

### Associating methods with data
- Method has a **receiver type** that it is associated with
- Use dot notation to call method

```go
type MyInt int

func (mi MyInt) Double() int {
	return int(mi * 2)
}

func main() {
	v := MyInt(3)
	fmt.Println(v.Double())
}
```

- Object v is an implicit argument to the method
  - Call by value

## 3 - Support for classes (2)
### Structs, again
- Struct types compose data fields

```go
type Point struct {
	x float64
	y float64
}
```

- Traditional feature of classes

### Structs with methods
- **Structs and methods** together allow arbitrary data and functions to be composed

```go
func (p Point) DistToOrigin() {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}

func main() {
	p1 := Point(3, 4)
	fmt.Println(p1.DistToOrigin())
}
```

## 4 - Encapsulation
### Controlling access
- Can define **public functions** to allow access to hidden data

```go
package data
var x int = 1
func PrintX() { fmt.Println(x) }

package main
import "data"
func main() {
	data.PrintX()
}
```

### Controlling access to structs
- Hide fields of structs by starting field name with a lower-case letter
- Define public methods which access hidden data

```go
package data

type Point struct {
	x float64
	y float64
}

func (p *Point) InitMe(xn, yn float64) {
	p.x = xn
	p.y = yn
}

func (p *Point) Scale(v float64) {
	p.x = p.x * v
	p.y = p.y * v
}

func (p *Point) PrintMe() {
	fmt.Println(p.x, p.y)
}

package main
func main() {
	var p data.Point
	p.InitMe(3, 4)
	p.Scale(2)
	p.PrintMe()
}
```

- Access to hidden fields only through public methods

## 5 - Point receivers
### Limitations of methods
- Receiver is passed implicitly as an argument to the method
- Method cannot modify the data inside the receiver
- Example: `OffsetX()` should increase x coordinate

```go
func main() {
	p1 := Point(3, 4)
	p1.OffsetX(5)
}
```

### Large receivers
- If receiver is large, lots of copying is required

```go
type Image [100][100]int

func main() {
	i1 := GrabImage()
	i1.BlurImage()
}
```

- 10,000 ints copied to BlurImage()

### Pointer receivers
```go
func (p *Point) OffsetX(v float64) {
	p.x = p.x + v
}
```

- Receiver can be a pointer to a type
- Call by reference, pointer is passed to the method

## 6 - Point receivers, referencing, dereferencing
### No need to dereference
```go
func (p *Point) OffsetX(v int) {
	p.x = p.x + v
}
```

- Point is referenced as `p`, not `*p`
- Dereferencing is automatic with `.` operator

### No need to reference
```go
func main() {
	p := Point(3, 4)
	p.OffsetX(5)
	fmt.Println(p.x)
}
```

### Using pointer receivers
- Good programming practice
  - All methods for a type have **pointer receivers**, or
  - All methods for a type have **non-pointer receivers**
- Mixing pointer/non-pointer receivers for a type will get confusing
  - Pointer receivers allows modification