package main

import (
	"testing"
)

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
	ans, err := readFile("./fixtures/test1.csv")
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

}

func TestParseQuiz(t *testing.T) {
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

	ans := parseQuiz(rows)
	if len(ans) != len(want) {
		t.Errorf("Expected %d got %d", len(want), len(ans))
	}
	for i := range ans {
		if ans[i] != want[i] {
			t.Errorf("Expected %v, got %v", want, ans)
		}
	}
}

func TestQuiz(t *testing.T) {
	prob := make([]problems, 2)
	prob[0] = problems{
		q: "5+5",
		a: "10",
	}
	prob[1] = problems{
		q: "7+5",
		a: "12",
	}

	quiz(prob, 5)
	// t.Log(ans)

}
