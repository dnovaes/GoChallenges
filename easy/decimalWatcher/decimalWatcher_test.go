package main

import "testing"

type Pair struct {
	input    string
	expected string
}

func TestDotFormat2Digits(t *testing.T) {
	tests := []Pair{
		Pair{"", ""},
		Pair{"0", "0.00"},
		Pair{"1", "1.00"},
		Pair{"10", "1.00"},
		Pair{"100", "1.00"},
		Pair{"1000", "10.00"},
		Pair{"100000", "1,000.00"},
		Pair{"1000000", "10,000.00"},
		Pair{"100000000", "1,000,000.00"},
		Pair{"100000000678", "1,000,000,006.78"},
		Pair{"1200000000678", "12,000,000,006.78"},
	}

	testFormats(tests, 2, t)
}

func TestDotFormat3Digits(t *testing.T) {
	tests := []Pair{
		Pair{"", ""},
		Pair{"0", "0.000"},
		Pair{"1", "1.000"},
		Pair{"10", "1.000"},
		Pair{"100", "1.000"},
		Pair{"1000", "1.000"},
		Pair{"100000", "100.000"},
		Pair{"1000000", "1,000.000"},
		Pair{"100000000", "100,000.000"},
		Pair{"100000000678", "100,000,000.678"},
		Pair{"1200000000678", "1,200,000,000.678"},
	}

	testFormats(tests, 3, t)
}

func TestEdgeCases(t *testing.T) {
	tests := []Pair{
		Pair{"1.10", "1.10"},
		Pair{"23.15", "23.15"},
	}
	testFormats(tests, 2, t)
}

func testFormats(cases []Pair, numDigits int, t *testing.T) {
	for i, test := range cases {
		result := decimalWatcher(test.input, numDigits)
		printErrorMessage(i, t, result, test.expected)
	}
}

func printErrorMessage(testNumber int, t *testing.T, result string, expected string) {
	if result != expected {
		t.Errorf("%d) result is '%s' instead of '%s'", testNumber, result, expected)
	}
}
