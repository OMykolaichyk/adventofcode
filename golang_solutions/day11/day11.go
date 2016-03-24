package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func IsContainIncreasingStraightOfThreeLetters(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i] == s[i-1]+1 && s[i] == s[i-2]+2 {
			return true
		}
	}
	return false
}

func IsContainIllegalLetters(s string) bool {
	return strings.ContainsAny(s, "iol")
}

func IsContainTwoPairs(s string) bool {
	countPairs := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			countPairs++
			i++
		}
	}
	return countPairs == 2
}

func IncrementStringAtPos(s string, pos int) string {
	runes := []rune(s)
	pos--
	runes[pos]++
	if runes[pos] > 'z' {
		runes[pos] = 'a'
		runes = []rune(IncrementStringAtPos(string(runes), pos))
	}
	return string(runes)
}

func IsValid(s string) bool {
	return IsContainTwoPairs(s) && !IsContainIllegalLetters(s) && IsContainIncreasingStraightOfThreeLetters(s)
}

func GetRidOfIllegalLetters(pwd string) string {
	for IsContainIllegalLetters(pwd) {
		if pos := strings.IndexAny(pwd, "iol"); pos != -1 {
			pwd = IncrementStringAtPos(pwd, pos+1)
			runes := []rune(pwd)
			for i := pos + 1; i < len(pwd); i++ {
				runes[i] = 'a'
			}
			pwd = string(runes)
		}
	}
	return pwd
}

func NextPassword(pwd string) string {
	for {
		if pwd = IncrementStringAtPos(pwd, len(pwd)); IsValid(pwd) {
			return pwd
		} else {
			if IsContainIllegalLetters(pwd) {
				pwd = GetRidOfIllegalLetters(pwd)
			}
		}
	}
}

func main() {
	inPtr := flag.String("i", "", "Input")
	flag.Parse()

	if *inPtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(NextPassword(*inPtr))
}
