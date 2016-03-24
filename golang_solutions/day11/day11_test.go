package main

import (
	"fmt"
	"testing"
)

func TestIsContainIncreasingStraightOfThreeLetters(t *testing.T) {
	if !IsContainIncreasingStraightOfThreeLetters("hijklmmn") {
		t.Error("hijklmmn does contain hij.")
	}
	if IsContainIncreasingStraightOfThreeLetters("abbceffg") {
		t.Error("abbceffg does not contain one increasing straight of at least three letters")
	}
	if IsContainIncreasingStraightOfThreeLetters("abbcegjk") {
		t.Error("abbcegjk does not contain one increasing straight of at least three letters")
	}
	if !IsContainIncreasingStraightOfThreeLetters("abcdefgh") {
		t.Error("abcdefgh does contain abc.")
	}
	if !IsContainIncreasingStraightOfThreeLetters("abcdffaa") {
		t.Error("abcdffaa does contain abc.")
	}
	if !IsContainIncreasingStraightOfThreeLetters("ghijklmn") {
		t.Error("ghijklmn does contain hij.")
	}
	if !IsContainIncreasingStraightOfThreeLetters("ghjaabcc") {
		t.Error("ghjaabcc does contain abc.")
	}
}

func TestIsContainIllegalLetters(t *testing.T) {
	//hijklmmn - fail
	if !IsContainIllegalLetters("hijklmmn") {
		t.Error("hijklmmn does contain illegal letter")
	}
	if IsContainIllegalLetters("abbceffg") {
		t.Error("abbceffg does not contain illegal letter")
	}
	if IsContainIllegalLetters("abbcegjk") {
		t.Error("abbcegjk does not contain illegal letter")
	}
	if IsContainIllegalLetters("abcdefgh") {
		t.Error("abcdefgh does not contain illegal letter")
	}
	if IsContainIllegalLetters("abcdffaa") {
		t.Error("abcdffaa does not contain illegal letter")
	}
	if !IsContainIllegalLetters("ghijklmn") {
		t.Error("ghijklmn does contain illegal letter")
	}
	if IsContainIllegalLetters("ghjaabcc") {
		t.Error("ghjaabcc does not contain illegal letter")
	}
}

func TestIsContainTwoPairs(t *testing.T) {
	if IsContainTwoPairs("hijklmmn") {
		t.Error("hijklmmn does not contain at least two different, non-overlapping pairs of letters")
	}
	if !IsContainTwoPairs("abbceffg") {
		t.Error("abbceffg does not contain at least two different, non-overlapping pairs of letters")
	}
	if IsContainTwoPairs("abbcegjk") {
		t.Error("abbcegjk does not contain at least two different, non-overlapping pairs of letters")
	}
	if IsContainTwoPairs("abcdefgh") {
		t.Error("abcdefgh does not contain at least two different, non-overlapping pairs of letters")
	}
	if !IsContainTwoPairs("abcdffaa") {
		t.Error("abcdffaa does not contain at least two different, non-overlapping pairs of letters")
	}
	if IsContainTwoPairs("ghijklmn") {
		t.Error("ghijklmn does not contain at least two different, non-overlapping pairs of letters")
	}
	if !IsContainTwoPairs("ghjaabcc") {
		t.Error("ghjaabcc does not contain at least two different, non-overlapping pairs of letters")
	}
}

func TestIsValid(t *testing.T) {
	if IsValid("hijklmmn") {
		t.Error("hijklmmn is invalid")
	}
	if IsValid("abbceffg") {
		t.Error("abbceffg is invalid")
	}
	if IsValid("abbcegjk") {
		t.Error("abbcegjk is invalid")
	}
	if IsValid("abcdefgh") {
		t.Error("abcdefgh is invalid")
	}
	if !IsValid("abcdffaa") {
		t.Error("abcdffaa is invalid")
	}
	if IsValid("ghijklmn") {
		t.Error("ghijklmn is invalid")
	}
	if !IsValid("ghjaabcc") {
		t.Error("ghjaabcc is invalid")
	}
}

func TestIncrementStringAtPos(t *testing.T) {
	if res := IncrementStringAtPos("aaa", 3); res != "aab" {
		t.Error(fmt.Sprintf("Result %s is not same as expected \"aaa\"", res))
	}
	if res := IncrementStringAtPos("aaa", 2); res != "aba" {
		t.Error(fmt.Sprintf("Result %s is not same as expected \"aba\"", res))
	}
	if res := IncrementStringAtPos("aaz", 3); res != "aba" {
		t.Error(fmt.Sprintf("Result %s is not same as expected \"aba\"", res))
	}
	if res := IncrementStringAtPos("azz", 3); res != "baa" {
		t.Error(fmt.Sprintf("Result %s is not same as expected \"baa\"", res))
	}
}

func TestGetRidOfIllegalLetters(t *testing.T) {
	legalLetters := GetRidOfIllegalLetters("hijklmmn")
	if IsContainIllegalLetters(legalLetters) {
		t.Error(fmt.Sprintf("Result %s contain illegal letters", legalLetters))
	}
	legalLetters = GetRidOfIllegalLetters("ghijklmn")
	if IsContainIllegalLetters(legalLetters) {
		t.Error(fmt.Sprintf("Result %s contain illegal letters", legalLetters))
	}
}

func TestNextPassword(t *testing.T) {
	if res := NextPassword("abcdefgh"); res != "abcdffaa" {
		t.Error(fmt.Sprintf("Result %s is not same as expected \"abcdffaa\"", res))
	}
	if res := NextPassword("ghijklmn"); res != "ghjaabcc" {
		t.Error(fmt.Sprintf("Result %s is not same as expected \"ghjaabcc\"", res))
	}
}
