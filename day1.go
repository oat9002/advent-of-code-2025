package main

import "fmt"

func Day1_1() {
	inputs, err := ReadInput("./inputs/day1.txt")
	if err != nil {
		return
	}

	fmt.Println(inputs)
}
