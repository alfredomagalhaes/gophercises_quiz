package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
done - read csv file
todo - create flag to customize csv file
todo - keep track of users points
todo - show the score to the user
*/

type Problem struct {
	Question string
	Answer   string
}

func readCsvFile(file string) []Problem {
	var questions []Problem

	fileContent, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("error reading csv file", err)
	}

	csvReader := csv.NewReader(bytes.NewReader(fileContent))
	for {
		var question Problem
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error reading csv file", err)
		}
		if len(record) < 2 {
			log.Fatal("csv format is not correct, need to have 2 columns")
		}

		question.Question = record[0]
		question.Answer = record[1]

		questions = append(questions, question)

	}

	return questions
}

func main() {

	var defaultCsvFile string = "problems.csv"
	var userCsv string
	var csvToRead string
	var totalPoints int
	var userScore int

	csvToRead = defaultCsvFile
	flag.StringVar(&userCsv, "csvFile", "", "use it to pass your own problems csv file")
	flag.Parse()

	if userCsv != "" {
		csvToRead = userCsv
	}

	questions := readCsvFile(csvToRead)
	totalPoints = len(questions)

	for _, question := range questions {
		fmt.Printf("answer the question, %s: ", question.Question)
		reader := bufio.NewReader(os.Stdin)
		// ReadString will block until the delimiter is entered
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.Replace(input, "\n", "", -1))
		if input == question.Answer {
			userScore++
		}
	}

	fmt.Printf("here is your score: %d of %d \n", userScore, totalPoints)
}
