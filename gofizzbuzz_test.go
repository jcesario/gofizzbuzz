package gofizzbuzz

import "testing"

func Testfbconv(t *testing.T) {
	// test int, expected result
	var m map[int]string
	m[3] = "fizz"
	m[14] = "14"
	m[15] = "fizzbuzz"

	for testint, expectedresult := range m {
		var result string
		result = fbconv(testint)
		if result != expectedresult {
			t.Error("test(%v): Expected %v, got %v ", testint, expectedresult, result)
		}
	}
}
