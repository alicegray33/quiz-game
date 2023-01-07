package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePtr := flag.String("filename", "problems.csv", "File name of the question/answer file in CSV format")
	flag.Parse()
	problemsFile := *filePtr

	csvFile := readFile(problemsFile)

	questions := getQuestions(csvFile)

	numQuestions := len(questions)
	printIntro(numQuestions)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	score := runGame(questions)
	fmt.Printf("You scored %v our of %v!\n", score, numQuestions)

}

func runGame(questions [][]string) uint {
	var score uint = 0
	var userAnswer string
	for _, qaPair := range questions {
		fmt.Println(qaPair[0])
		fmt.Scan(&userAnswer)
		if qaPair[1] == userAnswer {
			score += 1
		}
	}
	return score
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
	fmt.Println("Press enter to start")
}
