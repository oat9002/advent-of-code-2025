package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

func ReadInputFromWeb(path string) ([]string, error) {
	// TO DO
	response, err := http.Get(path)

	if err != nil {
		fmt.Printf("HTTP Get failed %s\n", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Read body failed %s\n", err)
		return nil, err
	}

	toReturn := strings.Split(string(body), "\n")

	return toReturn, nil
}
