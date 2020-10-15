package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("bad command")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])
	cmd := exec.Command(os.Args[2], os.arts)
	cmd.Run()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
