package main

import(
 "testing"
 "regexp"
)


func TestParseInstruction(t *testing.T) {
	if res := ParseInstruction("turn on"); res != TURNON {
		t.Error("Ivalid instruction:", res)
	}
	if res := ParseInstruction("turn off"); res != TURNOFF {
		t.Error("Ivalid instruction:", res)
	}
	if res := ParseInstruction("toggle"); res != TOGGLE {
		t.Error("Ivalid instruction:", res)
	}
}

func TestGetCoordinates(t *testing.T) {  
	var lineRegexp = regexp.MustCompile("^(turn on|turn off|toggle) (\\d+),(\\d+) through (\\d+),(\\d+)$")
	m := lineRegexp.FindStringSubmatch("toggle 0,0 through 0,0")
	coords := GetCoordinates(m)
	if coords.from.x != 0 || coords.from.y != 0 || coords.from.x != 0 || coords.from.x != 0  {
		t.Error("Invalid coordinates: ", coords)
	}
}
	
func TestParseLine(t *testing.T) {
	instr, coords := ParseLine("turn off 0,0 through 0,0")
	if instr != TURNOFF {
		t.Error("Invalid instruction:", instr)
	}
	if coords.from.x != 0 || coords.from.y != 0 || coords.from.x != 0 || coords.from.x != 0  {
		t.Error("Invalid coordinates:", coords)	
	}
	instr, coords = ParseLine("turn on 0,0 through 0,0")
	if instr != TURNON {
		t.Error("Invalid instruction:", instr)
	}
	if coords.from.x != 0 || coords.from.y != 0 || coords.from.x != 0 || coords.from.x != 0  {
		t.Error("Invalid coordinates:", coords)
	}
	instr, coords = ParseLine("toggle 0,0 through 0,0")
	if instr != TOGGLE {
		t.Error("Invalid instruction:", instr)
	}
	if coords.from.x != 0 || coords.from.y != 0 || coords.from.x != 0 || coords.from.x != 0  {
		t.Error("Invalid coordinates:", coords)	
	}
}

func TestTurnOnLights(t *testing.T) {
	coords := Coordinates{Point{0,0},Point{1,1}}
	res := make([]bool, 100000)
	TurnOnLights(&res, coords)	
	if !res[0] || !res[1] || !res[1000] || !res[1001] {
		t.Error("Invalid result:",res[0],res[1],res[1000],res[1001])
	}
} 

func TestTurnOffLights(t *testing.T) {
	coords := Coordinates{Point{0,0},Point{1,1}}
	res := make([]bool, 100000)
	TurnOffLights(&res, coords)
	if res[0] || res[1] || res[1000] || res[1001] {
		t.Error("Invalid result:",res[0],res[1],res[1000],res[1001])
	}
}

func TestToggleLights(t *testing.T) {
	coords := Coordinates{Point{0,0},Point{1,1}}
	res := make([]bool, 100000)
	ToggleLights(&res, coords)
	if !res[0] || !res[1] || !res[1000] || !res[1001] {
		t.Error("Invalid result:",res[0],res[1],res[1000],res[1001])
	}
}

func TestIncrementBrightness(t *testing.T) {
	coords := Coordinates{Point{0,0},Point{1,1}}
	res := make([]int, 100000)
	IncrementBrightness(&res, coords)
	if res[0] != 1 || res[1] != 1 || res[1000] != 1 || res[1001] != 1{
		t.Error("Invalid result:",res[0],res[1],res[1000],res[1001])
	}
}

func TestDecrementBrightness(t *testing.T) {
	coords := Coordinates{Point{0,0},Point{1,1}}
	res := make([]int, 100000)
	DecrementBrightness(&res, coords)
	if res[0] != 0 || res[1] != 0 || res[1000] != 0 || res[1001] != 0 {
		t.Error("Invalid result:",res[0],res[1],res[1000],res[1001])
	}
}

func TestIncreaseBrightness(t * testing.T) {
	coords := Coordinates{Point{0,0},Point{1,1}}
	res := make([]int, 1000000)
	IncreaseBrightness(&res, coords, 2)
	if  res[0] != 2 || res[1] != 2 || res[1000] != 2 || res[1001] != 2 {
		t.Error("Invalid result:",res[0],res[1],res[1000],res[1001])
	}
}

func TestCountTurnedOnLights(t *testing.T) {
	res := make([]bool, 1000000)
	for i:=0;i<1000;i++{
		res[i] = true
	}
	if res := CountTurnedOnLights(res); res != 1000 {
		t.Error("Invalid result:",res)
	}
}

func TestCountBrightness(t *testing.T) {
	res := make([]int, 1000000)
	for i:=0;i<1000;i++{
		res[i]++
	}
	if res := CountBrightness(res); res != 1000 {
		t.Error("Invalid result:",res)
	}
}

func TestProcessLinePart1(t *testing.T) {
	res := make([]bool, 1000000)
	ProcessLinePart1("turn on 0,0 through 999,999", &res)
	if c := CountTurnedOnLights(res); c != 1000000{
		t.Error("Invalid result:",c)
	}
	res = make([]bool, 1000000)
	ProcessLinePart1("toggle 0,0 through 999,0", &res)
	if c := CountTurnedOnLights(res); c != 1000{
		t.Error("Invalid result:",c)
	}
	res = make([]bool, 1000000)
	ProcessLinePart1("turn off 499,499 through 500,500", &res)
	if c := CountTurnedOnLights(res); c != 0 {
		t.Error("Invalid result:",c)
	}
	res = make([]bool, 1000000)
	res[499*1000+499] = true
	res[499*1000+500] = true
	res[500*1000+499] = true
	res[500*1000+500] = true
	ProcessLinePart1("turn off 499,499 through 500,500", &res)
	if c := CountTurnedOnLights(res); c != 0 {
		t.Error("Invalid result:",c)
	}
}

func TestProcessLinePart2(t *testing.T) {
	res := make([]int, 1000000)
	ProcessLinePart2("turn on 0,0 through 0,0", &res)
	if c := CountBrightness(res); c != 1{
		t.Error("Invalid result:",c)
	}
	res = make([]int, 1000000)
	ProcessLinePart2("toggle 0,0 through 999,999", &res)
	if c := CountBrightness(res); c != 2000000 {
		t.Error("Invalid result:",c)
	}
	res = make([]int, 1000000)
	ProcessLinePart2("turn off 0,0 through 999,999", &res)
	if c := CountBrightness(res); c != 0 {
		t.Error("Invalid result:",c)
	}
	res = make([]int, 1000000)
	ProcessLinePart2("turn on 0,0 through 999,999", &res)
	if c := CountBrightness(res); c != 1000000 {
		t.Error("Invalid result:",c)
	}
	res = make([]int, 1000000)
	ProcessLinePart2("turn on 0,0 through 999,999", &res)
	ProcessLinePart2("turn off 0,0 through 999,999", &res)
	ProcessLinePart2("toggle 0,0 through 999,999", &res)
	if c := CountBrightness(res); c != 2000000 {
		t.Error("Invalid result:",c)
	}
}