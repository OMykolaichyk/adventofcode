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

func TestIsNiceString(t *testing.T) {
	if !IsNiceString("ugknbfddgicrmopn") {
		t.Error("\"ugknbfddgicrmopn\" is nice, but result is false")
	}
	if !IsNiceString("aaa") {
		t.Error("\"aaa\" is nice, but result is false")
	}
	if IsNiceString("jchzalrnumimnmhp") {
		t.Error("\"jchzalrnumimnmhp\" is naughty, but result is true")
	}
	if IsNiceString("haegwjzuvuyypxyu") {
		t.Error("\"haegwjzuvuyypxyu\" is naughty, but result is true")
	}
	if IsNiceString("dvszwmarrgswjxmb") {
		t.Error("\"dvszwmarrgswjxmb\" is naughty, but result is true")
	}
}