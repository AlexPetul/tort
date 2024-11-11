package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Project struct {
	Name  string
	Owner string
}

func readInitial() ([]Project, error) {
	file, err := os.Open("./data/initial.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var result []Project
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		parts := strings.Split(name, "/")
		result = append(result, Project{
			Name:  parts[len(parts)-1],
			Owner: parts[len(parts)-2],
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return result, nil
}
