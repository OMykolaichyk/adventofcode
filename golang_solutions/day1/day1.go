package main

import(
		"os"
		"io/ioutil"
		"fmt"
		)

func main() {
	if len(os.Args) != 2 {
		panic("No input file provided")
	}

	input := os.Args[1]
	f, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	floor, pos := Process(string(f))
	fmt.Println("Floor = ", floor)
	fmt.Println("Pos = ", pos)
}

func Process(s string) (int, int) {
	var floor int = 0
	var flag bool = false
	var pos int = 1
	for i, c := range s {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
		if !flag && floor == -1 {
			flag = true
			pos += i
		} 
	}
	return floor, pos
}