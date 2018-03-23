package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	path string
)

func main() {
	parse()
	csvData := read()
    quizzes := getQuizzes(csvData)

	var score int
	for i, q := range quizzes {
		fmt.Printf("Problem #%v: %v = ", i+1, q.question)
		var input int
		fmt.Scanln(&input)
		if input == q.answer {
			score++
		}
	}
	fmt.Printf("You scored %v out of %v\n", score, len(quizzes))
}

func parse() {
	flag.StringVar(&path, "file", "problems.csv", "path to the csv file")
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
