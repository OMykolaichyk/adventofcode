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

	floor := Process(string(f))
	fmt.Println("Floor = ", floor)
}

func Process(s string) int {
	var floor int = 0
	for _, c := range s {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}
	return floor
}