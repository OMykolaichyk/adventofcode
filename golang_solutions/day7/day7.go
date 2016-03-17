package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Gate func(x, y uint16) uint16

type Wire struct {
	left  *Wire
	right *Wire
	gate    Gate
	val   interface{}
}

func (wire *Wire) value() (val uint16) {
	if wire != nil {
		if wire.val != nil {
			val = wire.val.(uint16)
		} else {
			val = wire.gate(wire.left.value(), wire.right.value())
			wire.val = val	
		}	
	}
	return
}

var gates = map[string]Gate {
	"":       noop,
	"NOT":    not,
	"AND":    and,
	"OR":     or,
	"LSHIFT": lshift,
	"RSHIFT": rshift,
}

type Wires struct {
	data map[string]*Wire
}

func NewWires() Wires {
	return Wires{data:make(map[string]*Wire)}
}

var wires = NewWires()

func (wires Wires) getWire(s string) (wire *Wire) {
	if s == "" {
		return 	nil
	}
	if isVal(s) {
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


var lineRegexp = regexp.MustCompile(
	"^(?:(?:([a-z]+|[0-9]+)) )?(?:(NOT|OR|AND|LSHIFT|RSHIFT) )?(?:([a-z]+|[0-9]+)) -> ([a-z]+)$")

func Process(s string) {
	list := lineRegexp.FindStringSubmatch(s)
	if list == nil {
		panic("Invalid cmd")
	}

	node := wires.getWire(list[4])

	gate, isPresent := gates[list[2]] 
	if!isPresent {
		panic("Unknown gate")
	} 
	node.gate = gate

	node.left = wires.getWire(list[1])
	node.right = wires.getWire(list[3])
}

func noop(x, y uint16) uint16 {
	return y
}

func not(x, y uint16) uint16 {
	return ^y
}

func and(x, y uint16) uint16 {
	return x & y
}

func lshift(x, y uint16) uint16 {
	return x << y
}

func rshift(x, y uint16) uint16 {
	return x >> y
}

func or(x, y uint16) uint16 {
	return x | y
}

func main() {
	if len(os.Args) != 2 {
		panic("No input file provided")
	}

	input := os.Args[1]

	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		Process(str)
	}
	fmt.Println("a", wires.getWire("a").value())
}

func ParseVal(s string) uint16 {
	num, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(num)
}

func isVal(s string) bool {
	if _, err := strconv.ParseUint(s, 10, 16); err != nil {
		return false
	}
	return true
}
