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
func partTwo(input []string) int {

	var multiplication int

	for i := 0; i < len(input); i++ {
		value_1, _ := strconv.Atoi(input[i])
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}
			value_2, _ := strconv.Atoi(input[j])
			if value_1+value_2 < 2020 {
				for k := 0; k < len(input); k++ {
					if k == j || k == i {
						continue
					}
					value_3, _ := strconv.Atoi(input[k])
					if value_3+value_2+value_1 == 2020 {
						multiplication = value_1 * value_2 * value_3
						break
					}
				}
			}
		}
	}

	return multiplication
}

func main() {
	fmt.Println(partTwo(readInput("./input.txt")))
}
