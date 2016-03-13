package main

import "testing"

func TestFindAnswerForSecretKey(t *testing.T) {
	if answ := FindAnswerForSecretKey("abcdef", "00000"); answ != 609043 {
		t.Error("")
	}
	if answ := FindAnswerForSecretKey("pqrstuv", "00000"); answ != 1048970 {
		t.Error("")
	}
}
