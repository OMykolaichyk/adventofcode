package main

import "testing"

func TestProcess(t *testing.T) {

	// (()) and ()() both result in floor 0.
	// ((( and (()(()( both result in floor 3.
	// ))((((( also results in floor 3.
	// ()) and ))( both result in floor -1 (the first basement level).
	// ))) and )())()) both result in floor -3.

	if fl, _ := Process("(())"); fl != 0 {
		t.Error("(()) result in floor 0")
	}

	if fl, _ := Process("()()"); fl != 0 {
		t.Error("()() result in floor 0")
	}

	if fl, _ := Process("((("); fl != 3 {
		t.Error("((( result in floor 3")
	}

	if fl, _ := Process("(()(()("); fl != 3 {
		t.Error("(()(()( result in floor 3")
	}

	if fl, _ := Process("))((((("); fl != 3 {
		t.Error("))((((( result in floor 3")
	}

	if fl, _ := Process("())"); fl != -1 {
		t.Error("()) result in floor -1")
	}

	if fl, _ := Process("))("); fl != -1 {
		t.Error("))( result in floor -1")
	}

	if fl, _ := Process(")))"); fl != -3 {
		t.Error("))) result in floor -3")
	}

	if fl, _ := Process(")())())"); fl != -3 {
		t.Error(")())()) result in floor -3")
	}

	// ) causes him to enter the basement at character position 1.
	// ()()) causes him to enter the basement at character position 5.

	if _, pos := Process(")"); pos != 1 {
		t.Error(") causes him to enter the basement at character position 1")
	}

	if _, pos := Process("()())"); pos != 5 {
		t.Error("()()) causes him to enter the basement at character position 5")
	}
}
