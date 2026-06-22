package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Printf("Enter a floating point number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	trunc := int(value)
	fmt.Println(trunc)
}