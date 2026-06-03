package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	var target int = rand.Intn(100) + 1
	// fmt.Println(target)
	for i := 9; i >= 0; i-- {

		fmt.Print("Enter a number: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if guess == target {
			fmt.Println("Good job! You guessed it!")
			return
		} else if guess < target {
			fmt.Println("Oops. Your guess was LOW.", i, "tries left!")
		} else {
			fmt.Println("Oops. Your guess was HIGH.", i, "tries left!")
		}
	}

	fmt.Println("Sorry. You didn't guess my number. It was:", target)
}