package main

import "testing"

func TestNewGates(t *testing.T) {
	gates := NewGates()
	assertNotPanic := func(t *testing.T, f func(string) Gate, s string) {
		defer func() {
			if r := recover(); r != nil {
				t.Error("The code panic")
			}
		}()
		f(s)
	}
	assertNotPanic(t, gates.GetGate, "")
	assertNotPanic(t, gates.GetGate, "NOT")
	assertNotPanic(t, gates.GetGate, "AND")
	assertNotPanic(t, gates.GetGate, "OR")
	assertNotPanic(t, gates.GetGate, "LSHIFT")
	assertNotPanic(t, gates.GetGate, "RSHIFT")

	assertPanic := func(t *testing.T, f func(string) Gate, s string) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("The code did not panic")
			}
		}()
		f(s)
	}

	assertPanic(t, gates.GetGate, "Unknown")
}

func TestGetGate(t *testing.T) {
	gates := NewGates()
	var x uint16 = 123
	var y uint16 = 456
	if gates.GetGate("")(x, y) != y {
		t.Error("Invalid result of NOOP")
	}
	if gates.GetGate("NOT")(x, y) != ^y {
		t.Error("Invalid result of NOT")
	}
	if gates.GetGate("AND")(x, y) != x&y {
		t.Error("Invalid result of AND")
	}
	if gates.GetGate("OR")(x, y) != x|y {
		t.Error("Invalid result of OR")
	}
	if gates.GetGate("LSHIFT")(x, y) != x<<y {
		t.Error("Invalid result of LSHIFT")
	}
	if gates.GetGate("RSHIFT")(y, x) != y>>x {
		t.Error("Invalid result of RSHIFT")
	}
}
