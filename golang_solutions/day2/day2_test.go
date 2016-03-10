package main

import "testing"

func unpack(args ...int) (int,int,int){
    	if len(args) != 3 {
    		panic(1)
    	}	
    	return args[0], args[1], args[2]
    }

func assertPanic(t *testing.T, f func(int, int, int), args ...int) {
	defer func() {
        if r := recover(); r == nil {
            t.Error("The code did not panic")
        } else {
        	if w, ok := r.(int); ok && w == 1 {
       			t.Error("Invalid number of args")        		
        	}
        }
    }()

    f(unpack(args...))
}

func TestStringToInt(t *testing.T) {
	if StringToInt("-1") != -1 {
		t.Error("Invalid string to int conversion result")
	}
	if StringToInt("1") != 1 {
		t.Error("Invalid string to int conversion result")
	}
	if StringToInt("0") != 0 {
		t.Error("Invalid string to int conversion result")
	}
}

func TestCheckIsValidBoxSidesSize(t *testing.T) {
	CheckIsValidBoxSidesSize(1, 1, 1)
	assertPanic(t, CheckIsValidBoxSidesSize, -1, 1, 1)
	assertPanic(t, CheckIsValidBoxSidesSize, 1, 0, 1)
}

func TestBoxSurfaceArea(t *testing.T) {
	if BoxSurfaceArea(2, 3, 4) != 52 {
		t.Error("Calculated box surface area is wrong")
	}
	if BoxSurfaceArea(1, 1, 10) != 42 {
		t.Error("Calculated box surface area is wrong")	
	}
	if BoxSurfaceArea(5, 5, 5) != 150 {
		t.Error("Calculated box surface area is wrong")
	}
}

func TestBoxSmallestSideArea(t *testing.T) {
	if BoxSmallestSideArea(2, 3, 4) != 6 {
		t.Error("Calculated smallest box side area is wrong")
	}
	if BoxSmallestSideArea(1, 1, 10) != 1 {
		t.Error("Calculated smallest box side area is wrong")	
	}
	if BoxSmallestSideArea(5, 5, 5) != 25 {
		t.Error("Calculated smallest box side area is wrong")	
	}
}

func TestParseInputData(t *testing.T) {
	if l,w,h := ParseInputData("2x3x4"); l != 2 || w != 3 || h != 4 {
		t.Error("Invalid result of parsing input data")	
	}
	if l,w,h := ParseInputData("1x1x10"); l != 1 || w != 1 || h != 10 {
		t.Error("Invalid result of parsing input data")	
	}
	if l,w,h := ParseInputData("5x5x5"); l != 5 || w != 5 || h != 5 {
		t.Error("Invalid result of parsing input data")	
	}
}

func TestWrappingPaper(t *testing.T) {
	if res := WrappingPaper("2x3x4"); res != 58{
		t.Error("Invalid result of calculating wrapping paper")
	}
	if res := WrappingPaper("1x1x10"); res != 43{
		t.Error("Invalid result of calculating wrapping paper")
	}
	if res := WrappingPaper("5x5x5"); res != 175{
		t.Error("Invalid result of calculating wrapping paper")
	}
}

func TestBowRibbon(t *testing.T) {
	if res := BowRibbon("2x3x4"); res != 24 {
		t.Error("Invalid result of calculating ribbon for bow")
	}
	if res := BowRibbon("1x1x10"); res != 10 {
		t.Error("Invalid result of calculating ribbon for bow")
	}
	if res := BowRibbon("5x5x5"); res != 125 {
		t.Error("Invalid result of calculating ribbon for bow")
	}
}

func TestWrappingRibbon(t *testing.T) {
	if res := WrappingRibbon("2x3x4"); res != 10{
		t.Error("Invalid result of calculating wrapping ribbon")
	}
	if res := WrappingRibbon("1x1x10"); res != 4{
		t.Error("Invalid result of calculating wrapping ribbon")
	}
	if res := WrappingRibbon("5x5x5"); res != 20{
		t.Error("Invalid result of calculating wrapping ribbon")
	}
}

func TestRibbon(t *testing.T) {
	if res := Ribbon("2x3x4"); res != 34{
		t.Error("Invalid result of calculating ribbon")
	}
	if res := Ribbon("1x1x10"); res != 14{
		t.Error("Invalid result of calculating ribbon")
	}
	if res := Ribbon("5x5x5"); res != 145{
		t.Error("Invalid result of calculating ribbon")
	}
}

func TestTwoMin(t *testing.T) {
	min1, min2 := TwoMin(5,0,1)
	if min1 != 0 || min2 != 1 {
		t.Error("Invalid result of TwoMin")
	}
	min1, min2 = TwoMin(5,4,6,2,1)
	if min1 != 1 || min2 != 2 {
		t.Error("Invalid result of TwoMin")
	}
}