package main

import (
	"fmt"
	"testing"
)

func TestGenerateAllPossibleRoutes(t *testing.T) {
	data := []string{
		"London", "Dublin", "Belfast",
	}
	expected := [][]string{
		{"Dublin", "London", "Belfast"},
		{"Dublin", "Belfast", "London"},
		{"London", "Dublin", "Belfast"},
		{"London", "Belfast", "Dublin"},
		{"Belfast", "Dublin", "London"},
		{"Belfast", "London", "Dublin"},
	}
	res := GenerateAllPossibleRoutes(data)
	if len(res) != len(expected) {
		t.Error("Sizes of result not as expected")
	}
	for _, r := range res {
		found := false
		for _, e := range expected {
			if fmt.Sprintf("%v", e) == fmt.Sprintf("%v", r) {
				found = true
			}
		}
		if !found {
			t.Error(fmt.Sprintf("%s not found in expected values", r))
		}
	}
}

func TestParseLine(t *testing.T) {
	data := "London to Dublin = 464"
	res := ParseLine(data)
	if res.distance != 464 {
		t.Error("Distance is not as expected")
	}
	for _, v := range res.locations {
		if !((v.to != "London" && v.from != "Dublin") || (v.to != "Dublin" && v.from != "London")) {
			fmt.Println(v.to, v.from)
			t.Error("Locations is not as expected")
		}
	}

}

func TestParse(t *testing.T) {
	data := []string{"London to Dublin = 464"}
	res := Parse(data)
	if dist, isPresent := res[LocationPair{"London", "Dublin"}]; isPresent {
		if dist != 464 {
			t.Error("")
		}
	} else {
		t.Error("")
	}
	if dist, isPresent := res[LocationPair{"Dublin", "London"}]; isPresent {
		if dist != 464 {
			t.Error("")
		}
	} else {
		t.Error("")
	}
}

func TestGetUniqueLocations(t *testing.T) {
	lines := []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}
	locations := Parse(lines)
	unique := GetUniqueLocations(locations)
	for i := range unique {
		for j := range unique {
			if i == j {
				continue
			}
			if unique[i] == unique[j] {
				t.Error("Not unique")
			}
		}
	}
}

func TestCalculateRouteDistance(t *testing.T) {
	lines := []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}
	expected := map[string]int{
		"[Dublin London Belfast]": 982,
		"[Dublin Belfast London]": 659,
		"[London Dublin Belfast]": 605,
		"[London Belfast Dublin]": 659,
		"[Belfast Dublin London]": 605,
		"[Belfast London Dublin]": 982,
	}
	locationsDistances := Parse(lines)
	unique := GetUniqueLocations(locationsDistances)
	allroutes := GenerateAllPossibleRoutes(unique)

	routesDistance := make(map[string]int)
	for _, v := range allroutes {
		routesDistance[fmt.Sprintf("%v", v)] = CalculateRouteDistance(v, locationsDistances)
	}

	for k, v := range routesDistance {
		if dist, isPresent := expected[k]; isPresent {
			if v != dist {
				t.Error(fmt.Sprintf("Distance is not as expected: ", dist))
			}
		} else {
			t.Error(fmt.Sprintf("Route is not found: %s", k))
		}
	}
}
