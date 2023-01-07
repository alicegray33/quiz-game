package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	problemsFile := "problems.csv"
	var userInput string
	var score uint = 0

	csvFile := readFile(problemsFile)

	questions := getQuestions(csvFile)

	printIntro(len(questions))

	for _, qaPair := range questions {
		fmt.Println(qaPair[0])
		fmt.Scan(&userInput)
		if qaPair[1] == userInput {
			score += 1
		}
	}

	fmt.Printf("You scored %v!", score)

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

func printIntro(numQuestions int) {
	fmt.Println("Welcome to the super-exciting quiz program!")
	fmt.Println("===========================================")
	fmt.Println("")
	fmt.Printf("You will be asked %v questions, then your score will be revealed at the end\n", numQuestions)
	fmt.Println("Press any key to start")
}
