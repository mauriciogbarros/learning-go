package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func (a Animal) Action(action string) {
	switch action {
	case "eat":
		a.Eat()

	case "move":
		a.Move()

	case "speak":
		a.Speak()

	default:
		fmt.Println("Invalid action")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

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
		animal_info := strings.Split(input, " ")
		if len(animal_info) != 2 {
			fmt.Println("Invalid input")
			continue
		}

		animal_info[0] = strings.ToLower(animal_info[0])
		animal_info[1] = strings.ToLower(animal_info[1])
		
		switch animal_info[0] {
		case "cow":
			cow.Action(animal_info[1])

		case "bird":
			bird.Action(animal_info[1])

		case "snake":
			snake.Action(animal_info[1])

		default:
			fmt.Println("Invalid animal")
		}
	}
}