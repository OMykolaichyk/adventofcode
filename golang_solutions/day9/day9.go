package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type LocationPair struct {
	from, to string
}

type LocationDistance struct {
	locations []LocationPair
	distance  int
}

func permute(prefix, s []string) [][]string {
	result := make([][]string, 0)
	if len(s) <= 1 {
		result = append(result, append([]string(nil), append(prefix, s...)...))
	} else {
		for i, v := range s {
			s = append(s[:i], s[i+1:]...)
			result = append(result, permute(append(prefix, v), s)...)
			s = append(s[:i], append([]string{v}, s[i:]...)...)
		}
	}
	return result
}

func GenerateAllPossibleRoutes(s []string) [][]string {
	return permute(make([]string, 0), s)
}

func ParseLine(s string) LocationDistance {
	lineRegexp := regexp.MustCompile("(.+) to (.+) = ([0-9]+)")
	m := lineRegexp.FindStringSubmatch(s)
	var res LocationDistance
	var err error
	if res.distance, err = strconv.Atoi(m[3]); err != nil {
		panic(err)
	}
	res.locations = []LocationPair{LocationPair{m[1], m[2]}, LocationPair{m[2], m[1]}}
	return res
}

func Parse(lines []string) map[LocationPair]int {
	result := make(map[LocationPair]int)
	for _, v := range lines {
		if len(v) != 0 {
			locationDistance := ParseLine(v)
			for _, v := range locationDistance.locations {
				result[v] = locationDistance.distance
			}
		}
	}
	return result
}

func GetUniqueLocations(locationDistances map[LocationPair]int) []string {
	uniqueLocations := make(map[string]bool)
	res := make([]string, 0)
	for k, _ := range locationDistances {
		if _, isPresent := uniqueLocations[k.from]; !isPresent {
			uniqueLocations[k.from] = false
			res = append(res, k.from)
		}
		if _, isPresent := uniqueLocations[k.to]; !isPresent {
			uniqueLocations[k.to] = false
			res = append(res, k.to)
		}
	}
	return res
}

func CalculateRouteDistance(route []string, distances map[LocationPair]int) int {
	var res int = 0
	for i := 1; i < len(route); i++ {
		res += distances[LocationPair{route[i-1], route[i]}]
	}
	return res
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

	locationDistances := Parse(strings.Split(string(f), "\n"))
	uniqueLocations := GetUniqueLocations(locationDistances)

	allRoutes := GenerateAllPossibleRoutes(uniqueLocations)

	routesDistance := make([]int, 0)
	for _, v := range allRoutes {
		routesDistance = append(routesDistance, CalculateRouteDistance(v, locationDistances))
	}
	sort.Ints(routesDistance)

	fmt.Println("Distance for the shortest route is:", routesDistance[0])
	fmt.Println("Distance for the longest route is:", routesDistance[len(routesDistance)-1])

}
