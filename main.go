package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

/*
done - read csv file
todo - create flag to customize csv file
todo - keep track of users points
todo - show the score to the user
*/

type problem struct {
	question string
	answer   string
}

func readCsvFile(file string) []problem {
	var questions []problem

	fileContent, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("error reading csv file", err)
	}

	csvReader := csv.NewReader(bytes.NewReader(fileContent))
	records, _ := csvReader.ReadAll()
	questions = make([]problem, len(records))
	for idx, row := range records {

		questions[idx] = problem{
			question: row[0],
			answer:   row[1],
		}

	}

	return questions
}

func main() {

	var csvToRead string
	var totalPoints int
	var userScore int
	var quizTimeOut int

	flag.StringVar(&csvToRead, "csvFile", "problems.csv", "use it to pass your own problems csv file")
	flag.IntVar(&quizTimeOut, "timer", 30, "use it to change the quiz time out timer, informed in seconds")
	flag.Parse()

	questions := readCsvFile(csvToRead)
	totalPoints = len(questions)

	timer := time.NewTimer(time.Duration(quizTimeOut) * time.Second)

problemLoop:
	for _, question := range questions {
		fmt.Printf("answer the question, %s: ", question.question)
		answerCh := make(chan string)
		go func() {
			reader := bufio.NewReader(os.Stdin)
			// ReadString will block until the delimiter is entered
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.Replace(input, "\n", "", -1))
			answerCh <- input
		}()

		select {
		case resp := <-answerCh:
			if resp == question.answer {
				userScore++
			}
		case <-timer.C:
			fmt.Println("")
			fmt.Println("times up...")
			break problemLoop
		}
	}

	fmt.Printf("here is your score: %d of %d \n", userScore, totalPoints)
}
