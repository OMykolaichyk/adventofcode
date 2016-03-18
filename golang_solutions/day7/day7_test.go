package main

import "testing"
import "strings"
import "fmt"

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
}

func TestProcess(t *testing.T) {
	var gates = NewGates()
	var wires = NewWires()

	for _, v := range testData {
		Process(v, gates, wires)
	}

	if res := wires.GetWire("d").Value(); res != 72 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 72))
	}
	if res := wires.GetWire("e").Value(); res != 507 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 507))
	}
	if res := wires.GetWire("f").Value(); res != 492 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 492))
	}
	if res := wires.GetWire("g").Value(); res != 114 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 114))
	}
	if res := wires.GetWire("h").Value(); res != 65412 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 65412))
	}
	if res := wires.GetWire("i").Value(); res != 65079 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 65079))
	}
	if res := wires.GetWire("x").Value(); res != 123 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 123))
	}
	if res := wires.GetWire("y").Value(); res != 456 {
		t.Error(fmt.Sprintf("Invalid result: %v.Expected: %v", res, 456))
	}
}