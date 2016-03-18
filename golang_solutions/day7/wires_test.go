package main

import "testing"

func TestIsVal(t *testing.T) {
	if !IsVal("123") {
		t.Error("\"123\" is a value")
	}
	if IsVal("var") {
		t.Error("\"var\" is not a value")
	}
	if IsVal("") {
		t.Error("\"\" is not a value")
	}
}

func TestParseVal(t *testing.T) {
	if ParseVal("123") != 123 {
		t.Error("result of ParseVal(\"123\") is not 123")
	}
	if ParseVal("0") != 0 {
		t.Error("result of ParseVal(\"\") is not 123")
	}

	assertPanic := func(t *testing.T, f func(string) uint16, s string) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("The code did not panic")
			}
		}()
		f(s)
	}
	assertPanic(t, ParseVal, "qw")

}

func TestGetWire(t *testing.T) {
	wires := NewWires()
	if wires.GetWire("") != nil {
		t.Error("")
	}
	a := wires.GetWire("a")
	if a == nil {
		t.Error("")
	}
	temp := wires.GetWire("a")
	if a != temp {
		t.Error("")
	}

	valWire := wires.GetWire("12")
	if valWire == nil {
		t.Error("")
	}

	temp = wires.GetWire("12")

	if valWire == temp {
		t.Error("")
	}
}

func TestWireValue(t *testing.T) {
	wires := NewWires()

	var wire *Wire

	if wire.Value() != 0 {
		t.Error("")
	}

	wire = wires.GetWire("12")
	if wire.Value() != 12 {
		t.Error("")
	}
}
