package main

import (
	"fmt"
	"strconv"
)

func Day1_1() {
	inputs, err := ReadInput("./inputs/day1.txt")
	if err != nil {
		return
	}
	position := 50
	numOfAtZero := 0

	for _, line := range inputs {
		rotationDirection := string(line[0])
		steps, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Printf("Atoi failed %s\n", err)
			break
		}

		//fmt.Printf("Current Position: %d, Direction: %s, Steps: %d\n", position, rotationDirection, steps)
		newPostion := rotate(rotationDirection, position, steps)
		//fmt.Printf("New Position: %d\n\n", newPostion)
		if newPostion == 0 {
			numOfAtZero++
		}
		position = newPostion
	}

	fmt.Printf("Total at 0: %d\n", numOfAtZero)
}

func Day1_2() {
	inputs, err := ReadInput("./inputs/day1.txt")
	if err != nil {
		return
	}
	position := 50
	numOfPassesAtZero := 0

	for _, line := range inputs {
		rotationDirection := string(line[0])
		steps, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Printf("Atoi failed %s\n", err)
			break
		}

		fmt.Printf("Current Position: %d, Direction: %s, Steps: %d\n", position, rotationDirection, steps)
		newPostion, n := rotate2(rotationDirection, position, steps)
		position = newPostion
		numOfPassesAtZero += n
		fmt.Printf("New Position: %d, passing 0: %d\n\n", newPostion, numOfPassesAtZero)

	}

	fmt.Printf("Total passing 0: %d\n", numOfPassesAtZero)
}

func rotate(direction string, currentPosition int, steps int) int {
	newPostion := 0

	switch direction {
	case "L":
		newPostion = currentPosition - steps
	case "R":
		newPostion = currentPosition + steps
	}

	for {
		if newPostion < 0 {
			newPostion = 100 + newPostion
		} else if newPostion > 99 {
			newPostion = newPostion - 100
		}

		if newPostion >= 0 && newPostion <= 99 {
			break
		}
	}

	return newPostion
}

func rotate2(direction string, currentPosition int, steps int) (int, int) {
	newPostion := 0
	numOfPassesAtZero := 0

	switch direction {
	case "L":
		newPostion = currentPosition - steps
	case "R":
		newPostion = currentPosition + steps
	}

	for {
		if newPostion < 0 {
			newPostion = 100 + newPostion
			numOfPassesAtZero++

		} else if newPostion > 99 {
			newPostion = newPostion - 100
			numOfPassesAtZero++

		}

		if newPostion >= 0 && newPostion <= 99 {
			break
		}
	}

	return newPostion, numOfPassesAtZero
}
