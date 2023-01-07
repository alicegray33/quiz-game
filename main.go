package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	problemsFile := "problems.csv"

	csvFile := readFile(problemsFile)

	questions := getQuestions(csvFile)

	fmt.Print(questions)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(problemsFile string) string {
	csvFile, err := os.ReadFile("./" + problemsFile)
	check(err)
	return string(csvFile)
}

func getQuestions(problemsFile string) [][]string {
	r := csv.NewReader(strings.NewReader(problemsFile))

	questions, err := r.ReadAll()
	check(err)

	return questions
}
