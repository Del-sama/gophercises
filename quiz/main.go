package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	const (
		defaultName      = "problems.csv"
		defaultTimeLimit = 30
	)
	fileName := flag.String("csv", defaultName, "csv file in the format: `question,answer`")
	timeLimit := flag.Int("limit", defaultTimeLimit, "Integer time limit")
	flag.Parse()

	res, err := quiz(*fileName, *timeLimit, os.Stdin)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
