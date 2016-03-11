package main 

import(
		"os"
		"io/ioutil"
		"fmt"
		)

type Coordinates struct {
	x, y int
}

func (pos *Coordinates) Move(cmd rune) {
	switch cmd {
		case '>':
			pos.x++
		case '<':
			pos.x--
		case '^':
			pos.y++
		case 'v':
			pos.y--
	}
}

func main() {
	if len(os.Args) != 2 {
		panic("No input file provided")
	}

	input := os.Args[1]
	f, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Number of houses visited by Santa is", CountHousesVisitedBySanta(string(f)))
	fmt.Println("Number of houses visited by Santa and Robo-Santa is", CountHousesVisitedBySantaAndRobot(string(f)))
}

func CountHousesVisitedBySanta(s string) int {
	housesCoords := make(map[Coordinates]struct{})
	currentSantaCoords := Coordinates{0, 0}
	housesCoords[currentSantaCoords] = struct{}{}

	for _, c := range s {
		currentSantaCoords.Move(c)
		housesCoords[currentSantaCoords] = struct{}{}
	}
	return len(housesCoords)
}

func CountHousesVisitedBySantaAndRobot(s string) int {
	housesCoords := make(map[Coordinates]struct{})
	currentSantaCoords := Coordinates{0, 0}
	currentRoboSantaCoords := Coordinates{0, 0}

	housesCoords[currentSantaCoords] = struct{}{}

	var current *Coordinates
	for i, c := range s {
		if i % 2 == 0 {
			current = &currentSantaCoords
		} else {
			current = &currentRoboSantaCoords
		}
		current.Move(c)
		housesCoords[*current] = struct{}{}
	}
	return len(housesCoords)
}