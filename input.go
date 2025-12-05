package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(path string) ([]string, error) {
	f, err := os.Open(path)

	if err == nil || len(f) == 0 {
		fmt.Println(err)

		return nil, err
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)

		return nil, err
	}

	return lines, nil
}
