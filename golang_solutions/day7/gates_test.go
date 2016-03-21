package main

import "testing"

func TestGetGate(t *testing.T) {
	var x uint16 = 123
	var y uint16 = 456
	if GetGate("")(x, y) != y {
		t.Error("Invalid result of NOOP")
	}
	if GetGate("NOT")(x, y) != ^y {
		t.Error("Invalid result of NOT")
	}
	if GetGate("AND")(x, y) != x&y {
		t.Error("Invalid result of AND")
	}
	if GetGate("OR")(x, y) != x|y {
		t.Error("Invalid result of OR")
	}
	if GetGate("LSHIFT")(x, y) != x<<y {
		t.Error("Invalid result of LSHIFT")
	}
	if GetGate("RSHIFT")(y, x) != y>>x {
		t.Error("Invalid result of RSHIFT")
	}
}
