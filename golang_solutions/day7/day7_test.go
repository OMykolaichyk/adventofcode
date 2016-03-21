package main

import (
	"fmt"
	"strings"
	"testing"
)

var testData = []string{
	"123 -> x", "456 -> y", "x AND y -> d",
	"x OR y -> e", "x LSHIFT 2 -> f", "y RSHIFT 2 -> g",
	"NOT x -> h", "NOT y -> i",
}

func TestParseLine(t *testing.T) {
	for _, v := range testData {
		parsed := ParseLine(v)
		list := strings.Split(v, " ")
		parsedLen := len(parsed)
		listLen := len(list)

		if listLen == 3 {
			if parsed[parsedLen-1] != list[2] && parsed[parsedLen-2] != list[0] {
				t.Error("Invalid parse result")
			}
		} else if listLen == 4 {
			if parsed[parsedLen-1] != list[3] && parsed[parsedLen-2] != list[1] && parsed[parsedLen-3] != list[0] {
				t.Error("Invalid parse result")
			}
		} else if listLen == 5 {
			if parsed[parsedLen-1] != list[4] && parsed[parsedLen-2] != list[2] && parsed[parsedLen-3] != list[1] && parsed[parsedLen-4] != list[0] {
				t.Error("Invalid parse result")
			}
		}
	}
	assertPanic := func(t *testing.T, f func(string) []string, s string) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("The code did not panic")
			}
		}()
		f(s)
	}

	assertPanic(t, ParseLine, "")
}

func TestProcess(t *testing.T) {
	var wires = NewWires()

	for _, v := range testData {
		Process(v, wires)
	}

	if res := wires.ValueOf("d"); res != 72 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 72))
	}
	if res := wires.ValueOf("e"); res != 507 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 507))
	}
	if res := wires.ValueOf("f"); res != 492 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 492))
	}
	if res := wires.ValueOf("g"); res != 114 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 114))
	}
	if res := wires.ValueOf("h"); res != 65412 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 65412))
	}
	if res := wires.ValueOf("i"); res != 65079 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 65079))
	}
	if res := wires.ValueOf("x"); res != 123 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 123))
	}
	if res := wires.ValueOf("y"); res != 456 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 456))
	}
}
