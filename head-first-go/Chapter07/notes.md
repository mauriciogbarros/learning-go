# Chapter 7. labeling data: Maps
## Counting votes

## Reading names from a file

## Counting names the hard way, with slices

## Maps
- A **map** is a collection where each value is accessed via a *key*.
- Kyes are an esay way to get data back out of your map.
- A map can use any type for keys
  - As long as values of that type can be compared using `==`
  - That includes numbers, strings, and more.
- The values all have to be of the same type
- The keys all have to be of the same type
- The keys don't have to be the same type as the values.
- To declare a map: `var myMap map[<key type>]<value type>`
- Just as with slices, declaring a map variable doesn't automatically create a map
  - You need to call the `make` function

```go
var ranks map[string]int // declare a map variable
ranks = make(map[string]int) // actually create the map

// or
ranks := make(map[stirng]int) // create a map and declare a variable to hold it.
```

- To assign and retrieve values:

```go
ranks["gold"] = 1
ranks["silver"] = 2
ranks["bronze"] = 3
fmt.Println(ranks["bronze"]) // 3
fmt.Println(ranks["gold"]) // 1
```

## Map literals
- Just as with arrays and slices, if you know keys and values that you want your map to have in advance, you can use a **map literal** to create it.

```go
myMap := map[string]float64{"a": 1.2, "b": 5.6}
```

- Empty map:

```go
emptyMap := map[string]float64{}
```

## Zero values within maps
- As with arrays and slices, if you access a map key that hasn't been assigned to, you will get a zero value back.
  - Maps with a `string` the zero value will be an empty string for example.

```go
numbers := make(map[string]int)
numbers["I've been assigned"] = 12
fmt.Printf("%#v\n", numbers["I've been assigned"]) // 12
fmt.Printf("%#v\n", numbers["I haven't been assigned"]) // 0
```

```go
counters := make(map[string]int)
counters["a"]++
counters["a"]++
counters["c"]++
fmt.Println(counters["a"], counters["b"], counters["c"]) // 2 0 1
```

## The zero value for map variable is nil
- As with slices, the zero vlaue for the map variable itself is `nil`.
- If no map exists to add new keys and values to, you will get a panic

```go
var nilMap map[int]string
fmt.Printf("%#v\n", nilMap) // map[int]string(nil)
nilMap[3] = "three" // panic: assignment to entry in nil map
```

- Before attempting to add keys and values, create a map using `make` or a map literal, and assing it to your map variable.

```go
var myMap map[int]string = make(map[int]string)
myMap[3] = "three"
fmt.Printf("%#v\n", myMap) // map[int]string{3:"three"}
```

## How to tell zero values apart from assigned values
```go
func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 85.5}
	grade := grades[name]
	if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}	
}

func main() {
	status("Alma") // Alma is failing!
	status("Carl") // Carl is failing!
}
```

- To address situations like this, accessing a map key optionally returns a second boolean value.
  - It will be `true` if the returned value has actually been assigned to the map, or
  - `false` if the returned value just represents the default zero value.

```go
counts := map[string]int{"a": 3, "b": 0}
var value int
var ok bool
value, ok = counters["a"]
fmt.Println(value, ok) // 3 true
value, ok = counters["b"]
fmt.Println(value, ok) // 0 true
value, ok = counters["c"]
fmt.Println(value, ok) // 0 false
```

## Removeing key/value pairs with the "delete" function
- Just pass the `delete` function two things:
  - The map you want to delete a key from
  - The key you want deleted.
  - That key and its corresponding value will be removed from the map.

## Using for...range loops with maps
```go
for key, value := range myMap {
	// Loop block here.
}
```

## The for...range loop handles maps in random order!
- A map is an *unordered* collection of keys and values.
- To go around this: two loops

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}
}
```