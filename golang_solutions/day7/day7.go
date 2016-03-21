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

func Process(s string, wires Wires) {
	list := ParseLine(s)
	wire := wires.GetWire(list[3])
	wire.gate = GetGate(list[1])
	wire.left = wires.GetWire(list[0])
	wire.right = wires.GetWire(list[2])
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

	var wires = NewWires()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		Process(str, wires)
	}
	a := wires.ValueOf("a")
	fmt.Println("Part1: a =", a)

	wires.SetValue("b", a)
	wires.ClearCache()

	fmt.Println("Part2: a =", wires.ValueOf("a"))
}
