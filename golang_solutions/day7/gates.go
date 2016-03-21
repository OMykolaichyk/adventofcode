package main

type Gate func(x, y uint16) uint16

func GetGate(s string) (gate Gate) {
	switch s {
	case "NOT":
		gate = func(x, y uint16) uint16 { return ^y }
	case "AND":
		gate = func(x, y uint16) uint16 { return x & y }
	case "OR":
		gate = func(x, y uint16) uint16 { return x | y }
	case "LSHIFT":
		gate = func(x, y uint16) uint16 { return x << y }
	case "RSHIFT":
		gate = func(x, y uint16) uint16 { return x >> y }
	default:
		gate = func(x, y uint16) uint16 { return y }
	}
	return
}
