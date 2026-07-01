package main

import (
	"fmt"
	"math"
)

func main() {
	a, v0, s0, err := GetParameters()
	if err != nil {
		fmt.Println(err)
	}
	DisplaceFn := GenDisplaceFn(a, v0, s0)
	t, err := GetTime()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Displacement: %0.2f\n", CalcDisplace(DisplaceFn, t))
}

func GetParameters() (float64, float64, float64, error) {
	var a, v0, s0 float64
	fmt.Println("Enter the parameters:")
	fmt.Print("Acceleration (a): ")
	if _, err := fmt.Scan(&a); err != nil {
		return 0, 0, 0, err
	}
	fmt.Print("Initial velocity (v0): ")
	if _, err := fmt.Scan(&v0); err != nil {
		return 0, 0, 0, err
	}
	fmt.Print("Initial displacement (s0): ")
	if _, err := fmt.Scan(&s0); err != nil {
		return 0, 0, 0, err
	}
	return a, v0, s0, nil
}

func GenDisplaceFn(a, v0, s0 float64) func (t float64) float64 {
	DisplaceFn := func(t float64) float64 {
		return 0.5 * a * math.Pow(t, 2) + v0 * t + s0
	}
	return DisplaceFn
}

func GetTime() (float64, error) {
	var t float64
	fmt.Print("Enter time (t): ")
	if _, err := fmt.Scan(&t); err != nil {
		return 0, err
	}
	return t, nil
}

func CalcDisplace(DisplaceFn func(float64) float64, t float64) float64 {
	return DisplaceFn(t)
}