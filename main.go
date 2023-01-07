package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var score uint = 0
var timeUp bool = false
var wg sync.WaitGroup

func main() {
	filePtr := flag.String("filename", "problems.csv", "File name of the question/answer file in CSV format")
	timerPtr := flag.Int("timer", 30, "Time limit to complete the quiz")
	randomFlatPTR := flag.Bool("random", false, "Boolean value on whether to randomize the questions")
	flag.Parse()
	problemsFile := *filePtr
	gameTime := *timerPtr
	randomize := *randomFlatPTR

	csvFile := readFile(problemsFile)

	questions := getQuestions(csvFile)
	numQuestions := len(questions)

	if randomize {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(numQuestions, func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
	}

	printIntro(numQuestions, gameTime)
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	wg.Add(1)
	go runGame(questions)
	go startTimer(gameTime)
	wg.Wait()

	if timeUp {
		fmt.Println("Times up!")
	}
	fmt.Printf("You scored %v our of %v!\n", score, numQuestions)

}

func startTimer(gameTime int) {
	defer wg.Done()
	time.Sleep(time.Duration(gameTime) * time.Second)
	timeUp = true
}

func runGame(questions [][]string) {
	defer wg.Done()
	var userAnswer string
	for _, qaPair := range questions {
		fmt.Println(qaPair[0])
		fmt.Scan(&userAnswer)
		if strings.EqualFold(qaPair[1], userAnswer) {
			score += 1
		}
	}
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

func printIntro(numQuestions, gameTime int) {
	fmt.Println("Welcome to the super-exciting quiz program!")
	fmt.Println("===========================================")
	fmt.Println("")
	fmt.Printf("You must answer %v questions in %v seconds, then your score will be revealed at the end\n", numQuestions, gameTime)
	fmt.Println("Press enter to start")
}
