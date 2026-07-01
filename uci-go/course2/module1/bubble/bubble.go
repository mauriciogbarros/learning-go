package main

import "fmt"

func main() {
	numbers, err := GetInput(10)
	if err != nil {
		fmt.Println(err)
		return
	}

	BubbleSort(numbers)
	fmt.Println(numbers)
}

func GetInput(amount int) ([]int, error) {
	var input = make([]int, amount)
	fmt.Printf("Enter %d numbers:", amount)
	for i := range amount {
		if _, err := fmt.Scan(&input[i]); err != nil {
			return nil, err
		}
	}

	return input, nil
}

func BubbleSort(numbers []int) {
	var swapped bool
	for i := 0; i < len(numbers) - 1; i++ {
		swapped = false
		for j := 0; j < len(numbers) - i - 1; j++ {
			if (numbers[j] > numbers[j + 1]){
				Swap(numbers, j)
				swapped = true
			}
		}
		if !swapped {
			break;
		}
	}
}

func Swap(slice []int, index int) {
	slice[index], slice[index + 1] = slice[index + 1], slice[index]
}