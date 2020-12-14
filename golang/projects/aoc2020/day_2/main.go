package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Password struct {
	maxValue     int
	minValue     int
	policyLetter string
	password     string
}

func readInput(inputFile string) []Password {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	var items []Password

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		minValue, _ := strconv.Atoi(strings.Split(parts[0], "-")[0])
		maxValue, _ := strconv.Atoi(strings.Split(parts[0], "-")[1])

		policyLetter := strings.TrimSuffix(parts[1], ":")

		password := parts[2]

		items = append(items, Password{
			minValue:     minValue,
			maxValue:     maxValue,
			policyLetter: policyLetter,
			password:     password,
		})
	}

	return items
}

func partOne(items []Password) int {
	validPasswords := 0

	for _, item := range items {
		count := strings.Count(item.password, item.policyLetter)
		if item.minValue <= count && item.maxValue >= count {
			validPasswords++
		}

	}

	return validPasswords
}

func partTwo(items []Password) int {
	validPasswords := 0

	for _, item := range items {
		counter := 0
		if string(item.password[item.minValue-1]) == item.policyLetter {
			counter++
		}

		if string(item.password[item.maxValue-1]) == item.policyLetter {
			counter++
		}

		if counter == 1 {
			validPasswords++
		}
	}

	return validPasswords
}

func main() {
	items := readInput("input.txt")
	fmt.Println("Passwords for Part one: ", partOne(items))
	fmt.Println("Passwords for Part two: ", partTwo(items))
}
