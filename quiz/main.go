package main

import (
	"flag"
	"log"
)

func main() {
	const (
		defaultName      = "problems.csv"
		defaultTimeLimit = 30
	)
	fileName := flag.String("csv", defaultName, "csv file in the format: `question,answer`")
	timeLimit := flag.Int("limit", defaultTimeLimit, "Integer time limit")
	flag.Parse()

	if !isValidFileName(*fileName) {
		log.Fatal("Invalid file name")
		return
	} else {
		arr, err := readFile(*fileName)
		if err != nil {
			log.Fatal(err)
			return
		}
		problems := parseQuiz(arr)
		quiz(problems, *timeLimit)
	}
}
