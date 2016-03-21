package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	numOfCharsOfCodeStrs := 0
	numOfCharsOfInMemoryStrs := 0
	numOfCharsOfEncodedStrs := 0

	for scanner.Scan() {
		bytesLine := scanner.Bytes()
		numOfCharsOfCodeStrs += len(bytesLine)

		inMemoryString, err := strconv.Unquote(string(bytesLine))
		if err != nil {
			fmt.Println("Error while Unquote string:", err)
		}
		numOfCharsOfInMemoryStrs += len(inMemoryString)

		encodedString := strconv.Quote(string(bytesLine))
		numOfCharsOfEncodedStrs += len(encodedString)
	}
	fmt.Println("Part1 result:", numOfCharsOfCodeStrs-numOfCharsOfInMemoryStrs)
	fmt.Println("Part2 result:", numOfCharsOfEncodedStrs-numOfCharsOfCodeStrs)
}
