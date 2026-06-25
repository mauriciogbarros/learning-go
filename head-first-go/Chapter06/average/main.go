package main

import (
	"fmt"
	"log"
)

func main() {
	numbers, err := GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	avg := Average(numbers...)
	fmt.Printf("Average: %0.2f\n", avg)
}