package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Date struct {
	year int
	month int
	day int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if d.month == 0 {
		return errors.New("Month is required to set the day")
	}
	var has31 = false
	for _, v := range [7]int {1, 3, 5, 7, 8, 10, 12} {
		if d.month == v {
			has31 = true
		}
	}
	if day < 1 || (d.month == 2 && day > 28) || (has31 && day > 31) || (!has31 && day > 30) {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func main() {
	var input string
	var exit = false
	for !exit {
		fmt.Print("Enter date (YYYY/MM/DD) or X to exit: ")
		length, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
		}
		if strings.ToLower(input) == "x" {
			exit = true
			fmt.Println("Exiting...")
			continue
		}
		if length != 1 {
			fmt.Println("Invalid date")
			continue
		}
		var dateData = strings.Split(input, "/")
		year, err := strconv.Atoi(dateData[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		month, err := strconv.Atoi(dateData[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		day, err := strconv.Atoi(dateData[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		date := Date{}
		err = date.SetYear(year)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = date.SetMonth(month)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = date.SetDay(day)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(date)
	}
	fmt.Println("Goodbye!")
}