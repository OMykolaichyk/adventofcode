package main

import (
	"bytes"
	"flag"
	"fmt"
	"strconv"
)

func count(r byte, s string) int {
	counter := 0
	for i := range s {
		if r == s[i] {
			counter++
		} else {
			break
		}
	}
	return counter
}

func lookAndSay(s string) string {
	var buffer bytes.Buffer
	i := 0
	for i < len(s) {
		c := count(s[i], s[i:])
		buffer.WriteString(strconv.Itoa(c))
		buffer.WriteString(string(s[i]))
		i += c
	}
	return buffer.String()
}

func process(s string, nIter int) int {
	for i := 0; i < nIter; i++ {
		s = lookAndSay(s)
	}
	return len(s)
}

func main() {
	inputPtr := flag.String("i", "", "Input")
	iterNumPtr := flag.Int("n", 0, "Number of iterations")

	flag.Parse()

	fmt.Println(process(*inputPtr, *iterNumPtr))
}
