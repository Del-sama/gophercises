package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
	"unicode"
)

type problems struct {
	q string
	a string
}

func quiz(problems []problems, timeLimit int) {
	correct := 0
	incorrect := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for idx, prob := range problems {
		fmt.Printf("Question number %d is %s? \n", idx+1, prob.q)
		select {
		case <-timer.C:
			fmt.Println("Ops! You ran out of time")
			fmt.Printf("out of %d questions, you got %d questions correct \n", len(problems), correct)
			return
		default:
			var response string
			fmt.Scanln(&response)
			if response == prob.a {
				fmt.Println("Genius! You got the answer right")
				correct++
			} else {
				fmt.Println("You failed that question but that doesn't make you a failure :)")
				incorrect++
			}
		}
	}
	fmt.Printf("out of %d questions, you got %d questions correct \n", len(problems), correct)
	return
}

func parseQuiz(arr [][]string) []problems {
	ret := make([]problems, len(arr))
	for i, row := range arr {
		ret[i] = problems{
			q: row[0],
			a: row[1],
		}
	}
	return ret
}

func readFile(fileName string) ([][]string, error) {
	problems, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file with filename %s \n", fileName)
		os.Exit(1)
	}
	header, err := bufio.NewReader(problems).ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	_, err = problems.Seek(int64(len(header)), io.SeekStart)
	if err != nil {
		return nil, err
	}
	updatedProblems := csv.NewReader(problems)

	rows, err := updatedProblems.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, err
}

func isValidFileName(fileName string) bool {
	if len(fileName) < 5 {
		return false
	}
	if fileName[len(fileName)-4:] != ".csv" {
		return false
	}
	newName := fileName[:len(fileName)-4]
	for _, val := range newName {
		if !unicode.IsLetter(val) {
			return false
		}
	}
	return true
}
