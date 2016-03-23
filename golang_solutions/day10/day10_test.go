package main

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	if res := count('1', "111"); res != 3 {
		t.Error(fmt.Sprintf("Result %v is invalid. Expected %v", res, 3))
	}
	if res := count('1', "121"); res != 1 {
		t.Error(fmt.Sprintf("Result %v is invalid. Expected %v", res, 1))
	}
	if res := count('2', "121"); res != 0 {
		t.Error(fmt.Sprintf("Result %v is invalid. Expected %v", res, 0))
	}
}

func TestLookAndSay(t *testing.T) {
	if res := lookAndSay("1"); res != "11" {
		t.Error(fmt.Sprintf("Result %v is invalid. Expected %v", res, "11"))
	}
	if res := lookAndSay("11"); res != "21" {
		t.Error(fmt.Sprintf("Result %v is invalid. Expected %v", res, "21"))
	}
	if res := lookAndSay("21"); res != "1211" {
		t.Error(fmt.Sprintf("Result %v is invalid. Expected %v", res, "1211"))
	}
	if res := lookAndSay("1211"); res != "111221" {
		t.Error(fmt.Sprintf("Result %v is invalid. Expected %v", res, "111221"))
	}
	if res := lookAndSay("111221"); res != "312211" {
		t.Error(fmt.Sprintf("Result %v is invalid. Expected %v", res, "312211"))
	}
	s := "1"
	for i := 0; i < 5; i++ {
		s = lookAndSay(s)
	}
	if s != "312211" {
		t.Error(fmt.Sprintf("Result %v is invalid. Expected %v", s, "312211"))
	}
}

func TestProcess(t *testing.T) {
	if res := process("1", 4); res != 6 {
		t.Error(fmt.Sprintf("Result %v is invalid. Expected %v", res, 6))
	}
}
