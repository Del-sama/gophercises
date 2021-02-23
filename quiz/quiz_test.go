package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func setupFile() {
	file, err := os.Create("test.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
		os.Exit(1)
	}
	fmt.Println("Successfully created file")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := [][]string{{"problem", "solution"}, {"5+5", "10"}}

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
			os.Exit(1)
		}
	}
	fmt.Println("Successfully updated file")
}

func teardownFile() {
	err := os.Remove("test.csv")
	if err != nil {
		log.Fatal("Cannot delete file", err)
		os.Exit(1)
	}
	fmt.Println("Successfully deleted file")
}

func TestInvalidFilename(t *testing.T) {
	var tests = []struct {
		filename string
		want     bool
	}{
		{"Problems.csv", true},
		{"problems.csv", true},
		{"Problems", false},
		{"pr0blems.csv", false},
		{"pro_blems.csv", false},
		{"problems..csv", false},
		{"problems.txt", false},
		{"test.csv", true},
		{".csv", false},
		{"t.csv", true},
	}
	for _, tt := range tests {
		ans := isValidFileName(tt.filename)
		if ans != tt.want {
			t.Errorf("got %t, want %t", ans, tt.want)
		}
	}
}

func TestReadFile(t *testing.T) {
	setupFile()
	ans, err := readFile("test.csv")
	if err != nil {
		t.Errorf("Got err %s", err)
	}
	want := make([][]string, 1)
	want[0] = []string{"5+5", "10"}
	if len(ans) != len(want) {
		t.Errorf("Expected %d got %d", len(want), len(ans))
	}
	for i := range ans[0] {
		if ans[0][i] != want[0][i] {
			t.Errorf("Expected %v, got %v", want, ans)
		}
	}
	teardownFile()
}

func TestParseProblems(t *testing.T) {
	rows := make([][]string, 2)
	rows[0] = []string{"5+5", "10"}
	rows[1] = []string{"7+5", "12"}

	want := make([]problems, 2)
	for i, row := range rows {
		want[i] = problems{
			q: row[0],
			a: row[1],
		}
	}

	ans := parseProblems(rows)
	if len(ans) != len(want) {
		t.Errorf("Expected %d got %d", len(want), len(ans))
	}
	for i := range ans {
		if ans[i] != want[i] {
			t.Errorf("Expected %v, got %v", want, ans)
		}
	}
}

func TestGetstdin(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("10\n"))
	res, err := getstdin(&stdin)
	if err != nil {
		t.Errorf("Got err %s ", err)
	}
	if strings.TrimSpace(res) != "10" {
		t.Errorf("Expected %s, got %s", "10", strings.TrimSpace(res))
	}
}

func TestQuiz(t *testing.T) {
	setupFile()
	file := "test.csv"
	limit := 5

	var stdin bytes.Buffer

	stdin.Write([]byte("10\n"))

	res, err := quiz(file, limit, &stdin)
	if err != nil {
		t.Errorf("Got err %s ", err)
	}
	want := "Out of 1 questions, you got 1 questions correct"

	if strings.TrimSpace(res) != want {
		t.Errorf("Expected %s, got %s", want, strings.TrimSpace(res))
	}

	stdin.Write([]byte("11\n"))
	res, err = quiz(file, limit, &stdin)
	if err != nil {
		t.Errorf("Got err %s ", err)
	}
	want = "Out of 1 questions, you got 0 questions correct"

	if strings.TrimSpace(res) != want {
		t.Errorf("Expected %s, got %s", want, strings.TrimSpace(res))
	}
	teardownFile()
}
