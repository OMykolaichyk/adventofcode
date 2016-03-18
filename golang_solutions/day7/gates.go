package main

import "fmt"

type Gate func(x, y uint16) uint16

type Gates struct {
	data map[string]Gate
}

func NewGates() Gates {
	return Gates{
		data: map[string]Gate{
			"":       func(x, y uint16) uint16 { return y },
			"NOT":    func(x, y uint16) uint16 { return ^y },
			"AND":    func(x, y uint16) uint16 { return x & y },
			"OR":     func(x, y uint16) uint16 { return x | y },
			"LSHIFT": func(x, y uint16) uint16 { return x << y },
			"RSHIFT": func(x, y uint16) uint16 { return x >> y },
		},
	}
}

func (gates Gates) GetGate(s string) Gate {
	gate, isPresent := gates.data[s]
	if !isPresent {
		panic(fmt.Sprintf("Unknown Gate \"%v\"", s))
	}
	return gate
}
