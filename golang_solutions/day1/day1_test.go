package main

import "testing"

// (()) and ()() both result in floor 0.
// ((( and (()(()( both result in floor 3.
// ))((((( also results in floor 3.
// ()) and ))( both result in floor -1 (the first basement level).
// ))) and )())()) both result in floor -3.

func TestProcess(t *testing.T) {
	expectedFloor := 0
	if Process("(())") != expectedFloor {
		t.Error("Result of Process func is not the same as expected")
	}
	if Process("()()") != expectedFloor {
		t.Error("Result of Process func is not the same as expected")	
	}
	expectedFloor = 3
	if Process("(((") != expectedFloor {
		t.Error("Result of Process func is not the same as expected")
	}
	if Process("(()(()(") != expectedFloor {
		t.Error("Result of Process func is not the same as expected")	
	}
	expectedFloor = 3 
	if Process("))(((((") != expectedFloor {
		t.Error("Result of Process func is not the same as expected")	
	}
	expectedFloor = -1
	if Process("())") != expectedFloor {
		t.Error("Result of Process func is not the same as expected")	
	}
	if Process("))(") != expectedFloor {
		t.Error("Result of Process func is not the same as expected")	
	}
	expectedFloor = -3
	if Process(")))") != expectedFloor {
		t.Error("Result of Process func is not the same as expected")	
	}
	if Process(")())())") != expectedFloor {
		t.Error("Result of Process func is not the same as expected")	
	}
}