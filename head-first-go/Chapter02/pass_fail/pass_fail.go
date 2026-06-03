// pass_fail reports whether a grade is passing or failing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")				// Prompt user to enter a grade.

	reader := bufio.NewReader(os.Stdin)	// Buffered reader that gets text
																			// from the keyboard.

	input, err := reader.ReadString('\n')		// Return everything the user has
																			// typed, up to where they pressed
																			// Enter key.

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	grade, err := strconv.ParseInt(input, 10, 32)

	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade == 10 {
		status = "perfect"
	} else if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}