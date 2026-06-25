package main

import "fmt"

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func main() {
	severalStrings("a", "b")
	severalStrings("a", "b", "c", "d", "e")
	severalStrings()
}