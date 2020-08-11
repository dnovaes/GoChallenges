package main

import "testing"

func TestSomething(t *testing.T) {
	result := alternate("beabeefeab")
	var expected int32 = 0
	PrintMessageError(1, t, result, expected)
}

func PrintMessageError(testNum int, t *testing.T, result int32, expected int32) {
	if result != expected {
		t.Errorf("%d) result is %d instead of %d", testNum, result, expected)
	}
}
