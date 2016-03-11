package main 

import(
		"fmt"
		"os"
		"strings"
		"bufio"
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
	var numOfNiceStrings int
	for scanner.Scan() {
		if IsNiceString(scanner.Text()) {
			numOfNiceStrings++
		}
	}
	fmt.Println("Number of nice string in input is", numOfNiceStrings)
}

func IsNiceString(s string) bool {
	return !DoesContainProbitedStrings(s) &&
			(CountVowels(s) >= 3) &&
			DoesContainAppearedSymbols(s)
}

func DoesContainProbitedStrings(s string) bool {
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