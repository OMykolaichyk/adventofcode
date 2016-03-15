package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Operation func(x, y uint16) uint16
type Type int

const (
	VAR Type = iota
	VAL
	NIL
)

type Arg struct {
	ThisType Type
	Var      string
	Val      uint16
}

type Args struct {
	x, y Arg
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

	fmt.Println("wait for a")
	time.Sleep(10 * time.Second)
	fmt.Println(<-signals["a"])
}

var signals = make(map[string]chan uint16)

func Process(s string) {
	args, op, dest := Parse(s)

	signals[dest] = make(chan uint16, 1)
	xChan := make(chan uint16, 1)
	yChan := make(chan uint16, 1)

	switch args.x.ThisType {
	case VAR:
		signals[args.x.Var] = xChan
	case VAL:
		xChan <- uint16(args.x.Val)
	case NIL:
		close(xChan)
	}

	switch args.y.ThisType {
	case VAR:
		signals[args.y.Var] = yChan
	case VAL:
		yChan <- uint16(args.y.Val)
	case NIL:
		close(yChan)
	}

	go func(x, y, dest chan uint16, op Operation) {
		var xVal uint16
		var yVal uint16
		// xFinish := false
		// yFinish := false
			select {
			case xVal, _ = <-x:
				// xFinish = true
			case yVal, _ = <-y:
				// yFinish = true
			}
		dest <- op(xVal, yVal)
	}(xChan, yChan, signals[dest], op)
}

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}

func ParseUint(s string) uint16 {
	num, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(num)
}

func Parse(s string) (args Args, op Operation, dest string) {
	list := strings.Split(s, " ")

	switch {
	case list[1] == "->":
		op = noop
		args.x = ParseArg1(list[0])
		args.y = ParseArg1("")
	case list[2] == "->":
		op = not
		args.x = ParseArg1("")
		args.y = ParseArg1(list[1])
	case list[3] == "->":
		op = getOp(list[1])
		args.x = ParseArg1(list[0])
		args.y = ParseArg1(list[2])
	}
	dest = list[len(list)-1]
	return
}

func ParseArg1(s string) (arg Arg) {
	if s == "" {
		arg.ThisType = NIL
		return
	}
	if num, err := strconv.ParseUint(s, 10, 16); err == nil {
		arg.ThisType = VAL
		arg.Val = uint16(num)
	}
	arg.ThisType = VAR
	arg.Var = s
	return
}

func getOp(s string) func(x, y uint16) uint16 {
	switch s {
	case "RSHIFT":
		return rshift
	case "LSHIFT":
		return lshift
	case "OR":
		return or
	case "AND":
		return and
	default:
		panic("Unkmowm Operatoin")
	}
}

func noop(x, y uint16) uint16 {
	return x
}

func not(x, y uint16) uint16 {
	return y
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
