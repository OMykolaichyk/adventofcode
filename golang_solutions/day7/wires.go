package main

import "strconv"

type Wire struct {
	left  *Wire
	right *Wire
	gate  Gate
	val   interface{}
}

func (wire *Wire) Value() (val uint16) {
	if wire != nil {
		if wire.val != nil {
			val = wire.val.(uint16)
		} else {
			val = wire.gate(wire.left.Value(), wire.right.Value())
			wire.val = val
		}
	}
	return
}

type Wires struct {
	data map[string]*Wire
}

func NewWires() Wires {
	return Wires{data: make(map[string]*Wire)}
}

func (wires Wires) GetWire(s string) (wire *Wire) {
	if s == "" {
		return nil
	}
	if IsVal(s) {
		wire = new(Wire)
		wire.val = ParseVal(s)
	} else {
		if w, isPresent := wires.data[s]; isPresent {
			wire = w
		} else {
			wire = new(Wire)
			wires.data[s] = wire
		}
	}
	return
}

func ParseVal(s string) uint16 {
	num, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(num)
}

func IsVal(s string) bool {
	if _, err := strconv.ParseUint(s, 10, 16); err != nil {
		return false
	}
	return true
}
