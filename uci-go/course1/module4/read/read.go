package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {

	var fileName string
	fmt.Print("Enter the name of a text file: ")
	fmt.Scanln(&fileName)
	fileName = strings.TrimSpace(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var names []name
	lines := bufio.NewScanner(file)
	for lines.Scan() {
		line := strings.Split(lines.Text()," ")
		names = append(names, name{line[0], line[1]})
	}
	if lines.Err() != nil {
		fmt.Println(lines.Err())
	}

	for _, name := range names {
		fmt.Printf("%s %s\n", name.fname, name.lname)
	}
}