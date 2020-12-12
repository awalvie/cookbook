package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput(file_path string) []string {
	var input []string

	f, _ := os.Open(file_path)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

// find the two numbers that add up to 2020
func partOne(input []string) int {

	var multiplication int

	for i := 0; i < len(input); i++ {
		value_1, _ := strconv.Atoi(input[i])
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}
			value_2, _ := strconv.Atoi(input[j])
			if value_1+value_2 == 2020 {
				multiplication = value_2 * value_1
				break
			}
		}
	}

	return multiplication
}

// find the product of three numbers that add up to 2020

func main() {
	fmt.Println(partOne(readInput("./input.txt")))
}
