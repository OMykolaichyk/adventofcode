package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var lineRegexp = regexp.MustCompile(
	"^(?:(?:([a-z]+|[0-9]+)) )?(?:(NOT|OR|AND|LSHIFT|RSHIFT) )?(?:([a-z]+|[0-9]+)) -> ([a-z]+)$")

func ParseLine(s string) []string {
	list := lineRegexp.FindStringSubmatch(s)
	if list == nil {
		panic("Invalid cmd")
	}
	return list[1:]
}

func Process(s string, gates Gates, wires Wires) {
	list := ParseLine(s)
	node := wires.GetWire(list[3])
	node.gate = gates.GetGate(list[1])
	node.left = wires.GetWire(list[0])
	node.right = wires.GetWire(list[2])
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

	var gates = NewGates()
	var wires = NewWires()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		Process(str, gates, wires)
	}
	fmt.Println("a", wires.GetWire("a").Value())
}
