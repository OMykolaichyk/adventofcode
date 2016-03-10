package main

import(
		"os"
		"fmt"
		"strings"
		"strconv"
		"bufio"
		"sort"
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

	var sum int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += WrapperPaper(scanner.Text())
	}
	fmt.Println("Total wrapper paper needed", sum)
}

func WrapperPaper(s string) (int) {
	l, w, h := ParseInputData(s)
	return BoxSurfaceArea(l, w, h) + BoxSmallestSideArea(l, w, h)
}

func ParseInputData(s string ) (l, w, h int) {
	list := strings.Split(s, "x")
	if len(list) != 3 {
		panic(fmt.Sprint("Parse failed: %s", s))
	}
	l = StringToInt(list[0])
	w = StringToInt(list[1])
	h = StringToInt(list[2])
	return
}

func StringToInt(s string) int {
	res, err:= strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

func CheckIsValidBoxSidesSize(l, w, h int) {
	if l < 1 || w < 1 || h < 1 {
		panic("Non positive numbers or zero is provided for calculatin surface area")
	}
}

func BoxSurfaceArea(l, w, h int) int {
	CheckIsValidBoxSidesSize(l, w, h)
	return (2*l*w + 2*w*h + 2*h*l)
}

func BoxSmallestSideArea(l, w, h int) int {
	CheckIsValidBoxSidesSize(l, w, h)
	lwh := []int{l, w ,h}
	sort.Ints(lwh)
	return lwh[0] * lwh[1]
}

