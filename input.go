package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(path string) ([]string, error) {
	f, err := os.Open(path)

	if err != nil {
		fmt.Printf("Open file failed %s\n", err)

		return nil, err
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		lines = append(lines, text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error, %s\n", err)

		return nil, err
	}

	return lines, nil
}
