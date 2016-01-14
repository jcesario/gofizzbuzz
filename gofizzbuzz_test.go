package gofizzbuzz

import "testing"
import "fmt"

type testpair struct {
	testint        int
	expectedresult string
}

var tests = []testpair{
	{3, "fizz"},
	{14, "14"},
	{15, "fizzbuzz"},
}

func TestFbconv(t *testing.T) {
	// test int, expected result
	for _, pair := range tests {
		fmt.Printf("test(%v): Expecting %v\n", pair.testint, pair.expectedresult)
		result := Fbconv(pair.testint)
		if result != pair.expectedresult {
			t.Error("test(%v): Expected %v, got %v\n", pair.testint, pair.expectedresult, result)
		}
	}
}
