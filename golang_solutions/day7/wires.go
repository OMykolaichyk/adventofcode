package main

import "strconv"

type Wire struct {
	name  string
	left  *Wire
	right *Wire
	gate  Gate
	val   uint16
}

func AssignValueToWire(wire *Wire, val uint16) {
	wire.left = CreateNamedWire(wire.name)
	wire.right = CreateValueWire(val)
	wire.gate = GetGate("")
}

func CreateNamedWire(name string) *Wire {
	wire := new(Wire)
	wire.name = name
	return wire
}

func CreateValueWire(val uint16) *Wire {
	wire := new(Wire)
	wire.val = val
	return wire
}

func (wire *Wire) Value(cache ...*map[string]uint16) (val uint16) {
	if wire != nil {
		if wire.left == nil && wire.right == nil {
			val = wire.val
		} else {
			if len(cache) > 0 {
				isPresent := false
				if val, isPresent = (*cache[0])[wire.name]; isPresent {
					return
				}
				val = wire.gate(wire.left.Value(cache[0]), wire.right.Value(cache[0]))
				(*cache[0])[wire.name] = val
			} else {
				val = wire.gate(wire.left.Value(), wire.right.Value())
			}

		}
	}
	return
}

type Wires struct {
	data  map[string]*Wire
	cache map[string]uint16
}

func NewWires() Wires {
	return Wires{data: make(map[string]*Wire), cache: make(map[string]uint16)}
}

func (wires Wires) GetWire(s string) (wire *Wire) {
	if s == "" {
		return nil
	}
	if IsVal(s) {
		wire = CreateValueWire(ParseVal(s))
	} else {
		if w, isPresent := wires.data[s]; isPresent {
			wire = w
		} else {
			wire = CreateNamedWire(s)
			wires.data[s] = wire
		}
	}
	return
}

func (wires *Wires) ClearCache() {
	wires.cache = make(map[string]uint16)
}

func (wires Wires) ValueOf(name string) (val uint16) {
	var isPresent bool
	if val, isPresent = wires.cache[name]; !isPresent {
		val = wires.GetWire(name).Value(&wires.cache)
	}
	return
}

func (wires Wires) SetValue(name string, val uint16) {
	AssignValueToWire(wires.GetWire(name), val)
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
