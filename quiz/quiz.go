package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"
	"unicode"
)

type problems struct {
	q string
	a string
}

func quiz(fileName string, timeLimit int) (string, error) {
	correct := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	fmt.Println(reflect.TypeOf(timer))

	if !isValidFileName(fileName) {
		return "", errors.New("Invalid file name")
	}
	arr, err := readFile(fileName)
	if err != nil {
		return "", err
	}
	problems := parseProblems(arr)
	for idx, prob := range problems {
		count, err := runQuiz(idx+1, prob.q, prob.a, timer)
		if err != nil {
			break
		}
		correct += count
	}
	return fmt.Sprintf(
		"Out of %d questions, you got %d questions correct \n",
		len(problems), correct), nil

}

func runQuiz(idx int, question string, answer string, timer *time.Timer) (int, error) {
	fmt.Printf("Question number %d is %s? \n", idx, question)
	responseCh := make(chan string)
	count := 0
	go func() {
		response := getstdin()
		responseCh <- response
	}()
	select {
	case <-timer.C:
		fmt.Println(`Ops! You ran out of time!`)
		return count, errors.New("out of time")
	case response := <-responseCh:
		if response == answer {
			fmt.Println("Genius! You got the answer right")
			count++
		} else {
			fmt.Println("You failed that question but that doesn't make you a failure :)")
		}
	}
	return count, nil
}

func getstdin() string {
	var response string
	fmt.Scanln(&response)
	return response
}

func parseProblems(arr [][]string) []problems {
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
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file with filename %s \n", fileName)
		os.Exit(1)
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	fmt.Println(rows[1:])
	return rows[1:], err
}

func isValidFileName(fileName string) bool {
	// Make sure filename does not contain special characters or numbers
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
