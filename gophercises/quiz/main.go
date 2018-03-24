package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	path    string
	timeout int
	score   int
)

func main() {
	parse()
	csvData := read()
	quizzes := getQuizzes(csvData)

	done := make(chan bool)
	go play(quizzes, done)
	select {
	case <-time.After((time.Duration(timeout)) * time.Second):
		fmt.Println("\nTime's up!")
		fmt.Printf("Your score is: %v!\n", score)
	case <-done:
		fmt.Printf("Your score is: %v!\n", score)
	}
}

func play(quizzes []quiz, done chan bool) {
	for i, q := range quizzes {
		fmt.Printf("Problem #%v: %v = ", i+1, q.question)
		var input int
		fmt.Scanln(&input)
		if input == q.answer {
			score++
		}
	}
	fmt.Printf("You scored %v out of %v\n", score, len(quizzes))
	done <- true
}

func parse() {
	flag.StringVar(&path, "file", "problems.csv", "path to the csv file")
	flag.IntVar(&timeout, "timeout", 100, "duration to run the quiz")
	flag.Parse()
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func read() string {
	absPath, err := filepath.Abs(path)
	checkError(err)

	data, err := ioutil.ReadFile(absPath)
	checkError(err)

	return string(data)
}

type quiz struct {
	question string
	answer   int
}

func getQuizzes(data string) []quiz {
	lines := strings.Split(data, "\n")

	var quizzes []quiz
	for _, line := range lines[:len(lines)-1] {
		split := strings.SplitN(line, ",", 2)

		ans, err := strconv.Atoi(split[1])
		checkError(err)

		quizzes = append(quizzes, quiz{
			question: split[0],
			answer:   ans,
		})
	}
	return quizzes
}
