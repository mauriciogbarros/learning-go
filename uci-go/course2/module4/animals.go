package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct {}

func (a cow) Eat() {
	fmt.Println("grass")
}

func (a cow) Move() {
	fmt.Println("walk")
}

func (a cow) Speak() {
	fmt.Println("moo")
}

type bird struct {}

func (a bird) Eat() {
	fmt.Println("worms")
}

func (a bird) Move() {
	fmt.Println("fly")
}

func (a bird) Speak() {
	fmt.Println("peep")
}

type snake struct {}

func (a snake) Eat() {
	fmt.Println("mice")
}

func (a snake) Move() {
	fmt.Println("slither")
}

func (a snake) Speak() {
	fmt.Println("hsss")
}

func NewAnimal(name string, animalType string, animals map[string]animal) error {
	switch animalType {
	case "cow":
		animals[name] = cow{}
	
	case "bird":
		animals[name] = bird{}

	case "snake":
		animals[name] = snake{}

	default:
		return errors.New("Invalid animal type.")
	}

	fmt.Println("Created it!")
	return nil
}

func QueryAnimal(name string, information string, animals map[string]animal) error {
	a, exists := animals[name]
	if !exists {
		return errors.New("Name not found.")
	}

	switch information {
	case "eat":
		a.Eat()

	case "move":
		a.Move()

	case "speak":
		a.Speak()

	default:
		return errors.New("Invalid action.")
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	animals := make(map[string]animal)

	for {
		fmt.Print(">")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if input ==  "\n" {
			return
		}

		input = strings.Trim(input, "\n")
		parameters := strings.Split(input, " ")
		if len(parameters) != 3 {
			fmt.Println("Invalid input")
			continue
		}

		command := strings.ToLower(parameters[0])
		name := strings.ToLower(parameters[1])
		what := strings.ToLower(parameters[2])
		
		switch command {
		case "newanimal":
			err := NewAnimal(name, what, animals)
			if err != nil {
				fmt.Println(err)
			}

		case "query":
			err := QueryAnimal(name, what, animals)
			if err != nil {
				fmt.Println(err)
			}

		default:
			fmt.Println("Invalid command.")
		}
	}
}