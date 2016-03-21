package main

import "testing"

func TestCreateNamedWire(t *testing.T) {
	wire := CreateNamedWire("a")
	if wire.name != "a" {
		t.Error("CreateNamedWire failed")
	}
}

func TestCreateValueWire(t *testing.T) {
	wire := CreateValueWire(123)
	if wire.val != 123 {
		t.Error("CreateValueWire failed")
	}
}

func TestAssignValueToWire(t *testing.T) {
	wire := CreateNamedWire("a")
	AssignValueToWire(wire, 123)
	if wire.right.val != 123 {
		t.Error("AssignValueToWire failed")
	}
}

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
		t.Error("GetWire creates wire with empty arg")
	}
	a := wires.GetWire("a")
	if a == nil {
		t.Error("Failed GetWire from empty Wires")
	}
	temp := wires.GetWire("a")
	if a != temp {
		t.Error("GetWire returned not the same wires on same name")
	}

	valWire := wires.GetWire("12")
	if valWire == nil {
		t.Error("Failed GetWire with value")
	}

	temp = wires.GetWire("12")

	if valWire == temp {
		t.Error("GetWire returned the same wires on same values")
	}
}

func TestWireValue(t *testing.T) {
	var wire *Wire

	if wire.Value() != 0 {
		t.Error("")
	}

	wire = CreateValueWire(12)
	if wire.Value() != 12 {
		t.Error("")
	}

	AssignValueToWire(wire, 12)
	if wire.Value() != 12 {
		t.Error("")
	}
}

func TestWireValueWithCaching(t *testing.T) {
	var wire *Wire
	cache := make(map[string]uint16)

	if wire.Value(&cache) != 0 {
		t.Error("Value of wire is invalid")
	}

	if len(cache) != 0 {
		t.Error("Cache must be emty")
	}

	wire = CreateValueWire(12)
	if wire.Value(&cache) != 12 {
		t.Error("Value of wire is invalid")
	}

	if len(cache) != 0 {
		t.Error("Cache must be empty")
	}

	wire.name = "a"
	AssignValueToWire(wire, 12)
	if wire.Value(&cache) != 12 {
		t.Error("Value of wire is invalid")
	}
	if len(cache) == 0 {
		t.Error("Cache must be not empty")
	}
	if cache["a"] != 12 {
		t.Error("Cached value of wire is invalid")
	}
}

func TestClearCache(t *testing.T) {
	wires := NewWires()
	a := wires.GetWire("a")
	AssignValueToWire(a, 123)
	_ = wires.ValueOf("a")
	wires.ClearCache()
	if len(wires.cache) != 0 {
		t.Error("Cache must be empty after ClearCache")
	}
}

func TestSetValue(t *testing.T) {
	wires := NewWires()
	a := wires.GetWire("a")
	wires.SetValue("a", 123)
	if a.Value() != 123 {
		t.Error()
	}
}
