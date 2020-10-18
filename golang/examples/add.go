package main

import (
	"fmt"
)

func add(x, y int) (sum int) {
	sum = x + y
	return
}

func main() {
	sum := add(5, 6)
	fmt.Print(sum)
}
