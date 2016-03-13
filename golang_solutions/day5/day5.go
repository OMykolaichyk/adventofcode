package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("No input file provided")
	}

	input := os.Args[1]

	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numOfNiceStringsPart1Rules int
	var numOfNiceStringsPart2Rules int
	for scanner.Scan() {
		str := scanner.Text()
		if IsNiceStringPart1Rules(str) {
			numOfNiceStringsPart1Rules++
		}
		if IsNiceStringPart2Rules(str) {
			numOfNiceStringsPart2Rules++
		}
	}
	fmt.Println("Number of nice string by part 1 rules in input is", numOfNiceStringsPart1Rules)
	fmt.Println("Number of nice string by part 2 rules in input is", numOfNiceStringsPart2Rules)
}

func IsNiceStringPart1Rules(s string) bool {
	return !DoesContainProhibitedStrings(s) &&
		(CountVowels(s) >= 3) &&
		DoesContainAppearedSymbols(s)
}

func IsNiceStringPart2Rules(s string) bool {
	return DoesContainAppearedSymbolsWithOneBetween(s) &&
		DoesContainAppearedSymbolPairs(s)
}

func DoesContainProhibitedStrings(s string) bool {
	return (strings.Index(s, "ab") != -1) ||
		(strings.Index(s, "cd") != -1) ||
		(strings.Index(s, "pq") != -1) ||
		(strings.Index(s, "xy") != -1)
}

func CountVowels(s string) int {
	return strings.Count(s, "a") +
		strings.Count(s, "e") +
		strings.Count(s, "i") +
		strings.Count(s, "o") +
		strings.Count(s, "u")
}

func DoesContainAppearedSymbols(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func DoesContainAppearedSymbolPairs(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if strings.Count(s, s[i:i+2]) > 1 {
			return true
		}
	}
	return false
}

func DoesContainAppearedSymbolsWithOneBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
