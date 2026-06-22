package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var sli = make([]int, 3)
	var input string
	var i int
	for input != "x" {
		fmt.Print("Enter a number or X to exit: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
			return
		}

		input = strings.ToLower(input)
		if input != "x" {
			number, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if i < 3 {
				sli[i] = number
				i++
			} else {
				sli = append(sli, number)
			}
			
			clone := slices.Clone(sli)
			slices.Sort(clone)
			fmt.Println(clone)
		}
	}
}