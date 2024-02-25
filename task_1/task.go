package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func getInput(inputTitle string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(inputTitle)
	input, _ := reader.ReadString('\n')

	return input
}

func validateInput(input string) bool {
	if !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(input) {
		fmt.Println("Invalid input:", input)
		return false
	}

	return true
}

func removeDuplicates(input []string) []string {
	prevStr := ""
	var result []string
	for i := 0; i < len(input); i++ {
		if input[i] != prevStr {
			result = append(result, input[i])
		}
		prevStr = input[i]
	}
	return result
}

func sortArray(input []string) []string {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	return input
}

func main() {
	seqOfWords := getInput(
		"Enter a a sequence of whitespace separated words: ",
	)

	if !validateInput(seqOfWords) {
		return
	}

	resultArray := removeDuplicates(
		sortArray(strings.Split(seqOfWords, " ")),
	)

	fmt.Println(strings.Join(resultArray, " "))
}
