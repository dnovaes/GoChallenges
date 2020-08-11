package main

import "testing"

const YES string = "YES"
const NO string = "NO"

func TestSameStartingPos(t *testing.T) {
	result := kangaroo2(0, 3, 0, 2)
	expected := NO
	printErrorMessage(0, t, result, expected)

	// this case returns YES and in hackerrank too, but in fact should return NO
	result = kangaroo2(0, 2, 0, 3)
	expected = YES
	printErrorMessage(1, t, result, expected)

	result = kangaroo2(0, 3, 0, 3)
	expected = NO
	printErrorMessage(2, t, result, expected)

	result = kangaroo2(0, 1, 0, 1)
	expected = NO
	printErrorMessage(3, t, result, expected)
}

func TestDiffStartingPos(t *testing.T) {
	result := kangaroo2(0, 3, 2, 2)
	expected := YES
	printErrorMessage(4, t, result, expected)

	result = kangaroo2(0, 3, 3, 2)
	expected = YES
	printErrorMessage(5, t, result, expected)

	result = kangaroo2(0, 5, 6, 2)
	expected = YES
	printErrorMessage(6, t, result, expected)

	result = kangaroo2(0, 2, 5, 3)
	expected = NO
	printErrorMessage(7, t, result, expected)

	result = kangaroo2(1, 5, 2, 2)
	expected = NO
	printErrorMessage(8, t, result, expected)
}

func printErrorMessage(testNumber int, t *testing.T, result string, expected string) {
	if result != expected {
		t.Errorf("%d) result is %s instead of %s", testNumber, result, expected)
	}
}
