package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var lineRegexp = regexp.MustCompile("^(turn on|turn off|toggle) (\\d+),(\\d+) through (\\d+),(\\d+)$")

type Point struct {
	x, y int
}

type Coordinates struct {
	from, to Point
}

type Instruction int

const (
	TURNON Instruction = iota
	TURNOFF
	TOGGLE
)

const (
	XSideSize = 1000
	YSideSize = 1000
)

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
	resPart1 := make([]bool, XSideSize*YSideSize)
	resPart2 := make([]int, XSideSize*YSideSize)
	for scanner.Scan() {
		str := scanner.Text()
		ProcessLinePart1(str, &resPart1)
		ProcessLinePart2(str, &resPart2)
	}
	fmt.Println(CountTurnedOnLights(resPart1))
	fmt.Println(CountBrightness(resPart2))
}

func ProcessLinePart1(s string, res *[]bool) {
	instr, coords := ParseLine(s)
	switch instr {
	case TURNON:
		TurnOnLights(res, coords)
	case TURNOFF:
		TurnOffLights(res, coords)
	case TOGGLE:
		ToggleLights(res, coords)
	}
}

func ProcessLinePart2(s string, res *[]int) {
	instr, coords := ParseLine(s)
	switch instr {
	case TURNON:
		IncreaseBrightness(res, coords, 1)
	case TURNOFF:
		DecrementBrightness(res, coords)
	case TOGGLE:
		IncreaseBrightness(res, coords, 2)
	}
}

func ParseLine(s string) (Instruction, Coordinates) {
	m := lineRegexp.FindStringSubmatch(s)
	if m == nil {
		panic(fmt.Errorf("invalid instruction %q", s))
	}
	instr := ParseInstruction(m[1])
	coords := GetCoordinates(m)
	return instr, coords
}

func ParseInstruction(s string) Instruction {
	if strings.Index(s, "turn on") != -1 {
		return TURNON
	} else if strings.Index(s, "turn off") != -1 {
		return TURNOFF
	} else if strings.Index(s, "toggle") != -1 {
		return TOGGLE
	} else {
		panic(fmt.Sprintf("Unknown instruction %s", s))
	}
}

func GetCoordinates(s []string) (coords Coordinates) {
	var err error
	coords.from.x, err = strconv.Atoi(s[2])
	if err != nil {
		panic(err)
	}
	coords.from.y, err = strconv.Atoi(s[3])
	if err != nil {
		panic(err)
	}
	coords.to.x, err = strconv.Atoi(s[4])
	if err != nil {
		panic(err)
	}
	coords.to.y, err = strconv.Atoi(s[5])
	if err != nil {
		panic(err)
	}
	return
}

func TurnOnLights(l *[]bool, coords Coordinates) {
	for i := coords.from.y; i <= coords.to.y; i++ {
		for j := coords.from.x; j <= coords.to.x; j++ {
			(*l)[i*YSideSize+j] = true
		}
	}
}

func TurnOffLights(l *[]bool, coords Coordinates) {
	for i := coords.from.y; i <= coords.to.y; i++ {
		for j := coords.from.x; j <= coords.to.x; j++ {
			(*l)[i*YSideSize+j] = false
		}
	}
}

func ToggleLights(l *[]bool, coords Coordinates) {
	for i := coords.from.y; i <= coords.to.y; i++ {
		for j := coords.from.x; j <= coords.to.x; j++ {
			(*l)[i*YSideSize+j] = !(*l)[i*YSideSize+j]
		}
	}
}

func IncrementBrightness(l *[]int, coords Coordinates) {
	for i := coords.from.y; i <= coords.to.y; i++ {
		for j := coords.from.x; j <= coords.to.x; j++ {
			(*l)[i*YSideSize+j]++
		}
	}
}

func DecrementBrightness(l *[]int, coords Coordinates) {
	for i := coords.from.y; i <= coords.to.y; i++ {
		for j := coords.from.x; j <= coords.to.x; j++ {
			if (*l)[i*YSideSize+j] > 0 {
				(*l)[i*YSideSize+j]--
			}
		}
	}
}

func IncreaseBrightness(l *[]int, coords Coordinates, incNum int) {
	for i := coords.from.y; i <= coords.to.y; i++ {
		for j := coords.from.x; j <= coords.to.x; j++ {
			(*l)[i*YSideSize+j] += incNum
		}
	}
}

func CountTurnedOnLights(l []bool) int {
	n := 0
	for _, v := range l {
		if v {
			n++
		}
	}
	return n
}

func CountBrightness(l []int) int {
	n := 0
	for _, v := range l {
		n += v
	}
	return n
}
