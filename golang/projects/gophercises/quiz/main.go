package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Problem struct {
	question string
	answer   string
}

func readCsv(filepath string) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Unable to read input file ", filepath, err)
	}
	defer f.Close()

	// reader for reading CSV
	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", filepath, err)
	}

	return records
}

func parseLines(records [][]string) []Problem {
	problems := make([]Problem, len(records))
	for i, record := range records {
		problems[i] = Problem{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}
	}
	return problems
}

func playQuiz(problems []Problem) int {

	correct := 0

	for i, problem := range problems {
		fmt.Printf("Question %d is: %s = ", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct = correct + 1
			fmt.Println("Correct Answer!")
		} else {
			fmt.Println("Oof, incorrect.")
		}
	}

	return correct

}
func main() {
	// options for cli
	problemsPtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	// limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	// usage message that is rendered when -h|--help is used
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "A simple quiz program written for the first Gophercises excerice\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// check if file exists
	records := readCsv(*problemsPtr)

	// parse records into question structs
	questions := parseLines(records)

	// let the quiz begin
	playQuiz(questions)

}
