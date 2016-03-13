package main

import "testing"

func TestCountHousesVisitedBySanta(t *testing.T) {
	if res := CountHousesVisitedBySanta(">"); res != 2 {
		t.Error("> delivers presents to 2 houses, but delivered to", res)
	}
	if res := CountHousesVisitedBySanta("^>v<"); res != 4 {
		t.Error("^>v< delivers presents to 4, but delivered to", res)
	}
	if res := CountHousesVisitedBySanta("^v^v^v^v^v"); res != 2 {
		t.Error("^v^v^v^v^v delivers presents to 4, but delivered to", res)
	}
}

func TestCountHousesVisitedBySantaAndRobot(t *testing.T) {
	if res := CountHousesVisitedBySantaAndRobot("^v"); res != 3 {
		t.Error("^v delivers presents to 3 houses, but delivered to", res)
	}
	if res := CountHousesVisitedBySantaAndRobot("^>v<"); res != 3 {
		t.Error("^>v< delivers presents to 3 houses, but delivered to", res)
	}
	if res := CountHousesVisitedBySantaAndRobot("^v^v^v^v^v"); res != 11 {
		t.Error("^v^v^v^v^v delivers presents to 11 houses, but delivered to", res)
	}
}
