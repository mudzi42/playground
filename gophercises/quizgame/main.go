package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
https://github.com/gophercises/quiz
*/

var (
	score          = 0
	totalQuestions int
)

type Quiz struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func startTimer(timeAllowed int) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press enter to begin timer.")

	reader.ReadString('\n')
	timer := time.NewTimer(time.Duration(timeAllowed) * time.Second)
	go func() {
		<-timer.C
		fmt.Println("\nTimes up!")
		calculateScore()
		os.Exit(0)
	}()
}

func calculateScore() {
	fmt.Println()
	fmt.Println("Quiz results:")
	fmt.Printf("Your final score: %d/%d\n", score, totalQuestions)
}

func shuffleQuestions(vals []Quiz) []Quiz {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Quiz, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

func readInQuiz(quizFile string) ([]Quiz, error) {
	csvFile, err := os.Open(quizFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %v", quizFile, err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var quiz []Quiz
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		quiz = append(quiz, Quiz{
			Question: line[0],
			Answer:   line[1],
		})
	}

	totalQuestions = len(quiz)
	return quiz, nil
}

func askQuestions(quiz []Quiz) error {
	var err error

	for i, q := range quiz {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%d) %s = ", i+1, q.Question)

		answer, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		formatedResponse := strings.ToLower(strings.TrimSpace(answer))
		formatedAnswer := strings.ToLower(strings.TrimSpace(q.Answer))

		fmt.Printf("Correct answer: %s - ", q.Answer)
		if formatedAnswer == formatedResponse {
			fmt.Println("Correct!!")
			score++
		} else {
			fmt.Println("Incorrect")
		}
	}

	return err
}

func main() {
	csvFile := flag.String("csvfile", "problems.csv", "question cvs file")
	timeAllowed := flag.Int("timeAllowed", 30, "time in seconds to finish quiz")
	shuffle := flag.Bool("shuffle", false, "shuttle questions")

	flag.Parse()
	quiz, err := readInQuiz(*csvFile)
	if err != nil {
		log.Fatal(err)
	}
	if *shuffle {
		quiz = shuffleQuestions(quiz)
	}

	if *csvFile == "bridgeofdeath.csv" {
		fmt.Printf("STOP!\nHe who would cross the Bridge of Death\nmust answer me these questions ere the other side he see.\n")
	} else {
		fmt.Println("Welcome to the Quiz Game")
	}

	fmt.Printf("You have %v seconds to finish.\n", *timeAllowed)
	startTimer(*timeAllowed)

	err = askQuestions(quiz)
	if err != nil {
		log.Fatal(err)
	}

	calculateScore()
}
