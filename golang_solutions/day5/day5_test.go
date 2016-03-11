package main

import "testing"

func TestDoesContainProhibitedStrings(t *testing.T) {
	if !DoesContainProhibitedStrings("haegwjzuvuyypxyu") {
		t.Error("\"haegwjzuvuyypxyu\" does contain prohibited string, but result is false")
	}
	if !DoesContainProhibitedStrings("acde") {
		t.Error("\"acde\" does contain prohibited string, but result is false")
	}
	if !DoesContainProhibitedStrings("apqe") {
		t.Error("\"apqe\" does contain prohibited string, but result is false")
	}
	if !DoesContainProhibitedStrings("acabe"){
		t.Error("\"acabe\" does contain prohibited string, but result is false")
	}
	if DoesContainProhibitedStrings("qweefdsdfe") {
		t.Error("\"qweefdsdfe\" doesn't contain prohibited string, but result is true")
	}
}

func TestCountVowels(t *testing.T) {
	if res := CountVowels("ugknbfddgicrmopn"); res != 3{
		t.Error("\"ugknbfddgicrmopn\" does contain 3 vowels, bur result is", res)
	}
	if res := CountVowels("aei"); res != 3 {
		t.Error("\"aei\" does contain 3 vowels, bur result is", res)
	}
	if res := CountVowels("xazegov"); res != 3 {
		t.Error("\"xazegov\" does contain 3 vowels, bur result is", res)
	}
	if res := CountVowels("aeiouaeiouaeiou"); res != 15 {
		t.Error("\"aeiouaeiouaeiou\" does contain 15 vowels, bur result is", res)
	}
	if res := CountVowels("dvszwmarrgswjxmb"); res != 1 {
		t.Error("\"dvszwmarrgswjxmb\" does contain 1 vowels, bur result is", res)
	}
}

func TestDoesContainAppearedSymbols(t *testing.T) {
	if !DoesContainAppearedSymbols("ugknbfddgicrmopn") {
		t.Error("\"ugknbfddgicrmopn\" contain appeared symbols, but result is false")
	}
	if !DoesContainAppearedSymbols("abcdde") {
		t.Error("\"abcdde\" does contain appeared symbols, but result is false")
	}
	if !DoesContainAppearedSymbols("aabbccdd") {
		t.Error("\"aabbccdd\" does contain appeared symbols, but result is false")
	}
	if DoesContainAppearedSymbols("jchzalrnumimnmhp") {
		t.Error("\"jchzalrnumimnmhp\" doesn't contain appeared symbols, but result is true")
	}
}

func TestIsNiceStringPart1Rules(t *testing.T) {
	if !IsNiceStringPart1Rules("ugknbfddgicrmopn") {
		t.Error("\"ugknbfddgicrmopn\" is nice, but result is false")
	}
	if !IsNiceStringPart1Rules("aaa") {
		t.Error("\"aaa\" is nice, but result is false")
	}
	if IsNiceStringPart1Rules("jchzalrnumimnmhp") {
		t.Error("\"jchzalrnumimnmhp\" is naughty, but result is true")
	}
	if IsNiceStringPart1Rules("haegwjzuvuyypxyu") {
		t.Error("\"haegwjzuvuyypxyu\" is naughty, but result is true")
	}
	if IsNiceStringPart1Rules("dvszwmarrgswjxmb") {
		t.Error("\"dvszwmarrgswjxmb\" is naughty, but result is true")
	}
}

func TestDoesContainAppearedSymbolPairs(t *testing.T) {
	if !DoesContainAppearedSymbolPairs("xyxy") {
		t.Error("\"xyxy\" does contain a pair of any two letters that appears at least twice in the string without overlapping, but result is false")
	}
	if !DoesContainAppearedSymbolPairs("aabcdefgaa") {
		t.Error("\"aabcdefgaa\" does contain a pair of any two letters that appears at least twice in the string without overlapping, but result is false")
	}
	if !DoesContainAppearedSymbolPairs("qjhvhtzxzqqjkmpb") {
		t.Error("\"qjhvhtzxzqqjkmpb\" does contain a pair of any two letters that appears at least twice in the string without overlapping, but result is false")
	}
	if !DoesContainAppearedSymbolPairs("xxyxx") {
		t.Error("\"xxyxx\" does contain a pair of any two letters that appears at least twice in the string without overlapping, but result is false")
	}
	if !DoesContainAppearedSymbolPairs("uurcxstgmygtbstg") {
		t.Error("\"uurcxstgmygtbstg\" does contain a pair of any two letters that appears at least twice in the string without overlapping, but result is false")
	}
	if DoesContainAppearedSymbolPairs("ieodomkazucvgmuy") {
		t.Error("\"ieodomkazucvgmuy\" doesn't contain a pair of any two letters that appears at least twice in the string without overlapping, but result is true")
	}
	if DoesContainAppearedSymbolPairs("aaa") {
		t.Error("\"aaa\" does contain a pair of any two letters that appears at least twice in the string with overlapping!, but result is true")
	}
}


//does contain at least one letter which repeats with exactly one letter between them
func TestDoesContainAppearedSymbolsWithOneBetween(t *testing.T) {
	if !DoesContainAppearedSymbolsWithOneBetween("xyx") {
		t.Error("\"xyx\" does contain at least one letter which repeats with exactly one letter between them, but result is false")
	}
	if !DoesContainAppearedSymbolsWithOneBetween("aaa") {
		t.Error("\"aaa\" does contain at least one letter which repeats with exactly one letter between them, but result is false")
	}
	if !DoesContainAppearedSymbolsWithOneBetween("abcdefeghi") {
		t.Error("\"abcdefeghi\" does contain at least one letter which repeats with exactly one letter between them, but result is false")
	}
	if !DoesContainAppearedSymbolsWithOneBetween("qjhvhtzxzqqjkmpb") {
		t.Error("\"qjhvhtzxzqqjkmpb\" does contain at least one letter which repeats with exactly one letter between them, but result is false")
	}
	if !DoesContainAppearedSymbolsWithOneBetween("xxyxx") {
		t.Error("\"xxyxx\" does contain at least one letter which repeats with exactly one letter between them, but result is false")
	}
	if DoesContainAppearedSymbolsWithOneBetween("uurcxstgmygtbstg") {
		t.Error("\"uurcxstgmygtbstg\" does contain at least one letter which repeats with exactly one letter between them, but result is false")
	}
	if !DoesContainAppearedSymbolsWithOneBetween("ieodomkazucvgmuy") {
		t.Error("\"ieodomkazucvgmuy\" does contain at least one letter which repeats with exactly one letter between them, but result is true")
	}
}

func TestIsNiceStringPart2Rules(t *testing.T) {
	if !IsNiceStringPart2Rules("qjhvhtzxzqqjkmpb") {
		t.Error("\"qjhvhtzxzqqjkmpb\" is nice, but result is false")
	}
	if !IsNiceStringPart2Rules("xxyxx") {
		t.Error("\"xxyxx\" is nice, but result is false")
	}
	if IsNiceStringPart2Rules("uurcxstgmygtbstg") {
		t.Error("\"uurcxstgmygtbstg\" is naughty, but result is true")
	}
	if IsNiceStringPart2Rules("ieodomkazucvgmuy") {
		t.Error("\"ieodomkazucvgmuy\" is naughty, but result is true")
	}
}

