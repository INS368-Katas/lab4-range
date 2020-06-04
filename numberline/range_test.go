package numberline

import (
	"testing"
)

var rangeTest Range

func TestNewRange(t *testing.T) {

	var tests = []struct{
		input			string
		expected	Range
		err				error
	}{
		{"(2,5]", Range{3,5}, nil},
		{"[-2,15)", Range{-2,14}, nil},
		{"-2,15", Range{}, errInvalidRange},
		{"(-f,15)", Range{}, errInvalidValue},
	}

	for _,test := range tests {
		output, err := rangeTest.NewRange(test.input)

		if output != (Range{}) && output != test.expected {
			t.Errorf("Testing Failed: \"%v\" inputted, %v expected, received: %v", test.input, test.expected, output)
		}

		if err != test.err {
			t.Errorf("Testing Failed: \"%v\" inputted, \"%v\" expected, received: \"%v\"", test.input, test.err, err)
		}

	}
}

func TestContains(t *testing.T) {
	rangeTest = Range{3,5}

	var tests = []struct{
		input			[]int
		expected	bool
	}{
		{[]int{3, 2, 5, 7}, false},
		{[]int{3, 3, 5, 4, 5}, true},
		{[]int{3, 5, 6}, false},
		{[]int{4}, true},
	}

	for _,test := range tests {
		output := rangeTest.Contains(test.input...)
		if output != test.expected {
			t.Errorf("Testing Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

func TestDoesNotContain(t *testing.T) {
	rangeTest = Range{3,5}

	var tests = []struct{
		input			[]int
		expected	bool
	}{
		{[]int{-1, 2, 7, 5}, true},
		{[]int{3, 5}, false},
		{[]int{1, 6}, true},
		{[]int{5}, false},
	}

	for _,test := range tests {
		output := rangeTest.DoesNotContain(test.input...)
		if output != test.expected {
			t.Errorf("Testing Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

func TestGetAllPoints(t *testing.T) {

	var tests = []struct{
		rangeTest		Range
		expected		[]int
	}{
		{Range{3, 5}, []int{3, 4, 5}},
		{Range{0, 7}, []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{Range{5, 5}, []int{5}},
		{Range{0, 0}, []int{0}},
	}

	for _,test := range tests {
		output := test.rangeTest.GetAllPoints()
		if !compare(output, test.expected) {
			t.Errorf("Testing Failed: %v expected, received: %v", test.expected, output)
		}
	}
}

func TestContainsRange(t *testing.T) {

	var tests = []struct{
		input			string
		rangeTest	Range
		expected	bool
	}{
		{"(2,6]", Range{3,5}, false},
		{"(2,4]", Range{3,5}, true},
		{"[0,5)", Range{1,10}, false},
		{"[2,10]", Range{1,10}, true},
	}

	for _,test := range tests {
		output := test.rangeTest.ContainsRange(test.input)
		if output != test.expected {
			t.Errorf("Testing Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

func TestDoesNotContainRange(t *testing.T) {

	var tests = []struct{
		input			string
		rangeTest	Range
		expected	bool
	}{
		{"(2,6]", Range{3,5}, true},
		{"(2,4]", Range{3,5}, false},
		{"(0,5)", Range{1,10}, false},
		{"(-1,7)", Range{1,10}, true},
	}

	for _,test := range tests {
		output := test.rangeTest.DoesNotContainRange(test.input)
		if output != test.expected {
			t.Errorf("Testing Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

func TestGetEndPoints(t *testing.T) {

	var tests = []struct{
		rangeTest		Range
		lowerBound	int
		upperBound	int		
	}{
		{Range{3, 5}, 3, 5},
		{Range{1, 7}, 1, 7},
		{Range{0,0}, 0, 0},
		{Range{6,6}, 6, 6},
	}

	for _,test := range tests {
		lower, upper := test.rangeTest.GetEndPoints()

		if lower != test.lowerBound && upper != test.upperBound {
			t.Errorf("Testing Failed: (%v,%v) expected, received: (%v,%v)", test.lowerBound, test.upperBound, lower, upper)
		}
	}
}

func TestOverlapsRange(t *testing.T) {
	rangeTest = Range{3, 5}

	var tests = []struct{
		input			string
		expected	bool
	}{
		{"(4,9]", true},
		{"[-1,3]", true},
		{"(5,8)", false},
		{"(-2,2]", false},
	}

	for _,test := range tests {
		output := rangeTest.OverlapsRange(test.input)
		if output != test.expected {
			t.Errorf("Testing Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

func TestEquals(t *testing.T) {
	rangeTest = Range{3, 5}

	var tests = []struct{
		input			string
		expected	bool
	}{
		{"(2,5]", true},
		{"[3,6)", true},
		{"(3,5]", false},
		{"(-1,3]", false},
	}

	for _,test := range tests {
		output := rangeTest.Equals(test.input)
		if output != test.expected {
			t.Errorf("Testing Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

func TestNotEquals(t *testing.T) {
	rangeTest = Range{3, 5}

	var tests = []struct{
		input			string
		expected	bool
	}{
		{"(2,5]", false},
		{"(3,5]", true},
		{"(-1,3]", true},
		{"[3,6)", false},
	}

	for _,test := range tests {
		output := rangeTest.NotEquals(test.input)
		if output != test.expected {
			t.Errorf("Testing Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

func compare(output []int, expected []int) bool {
	for index := range expected {
		if output[index] != expected[index] {
			return false
		}
	}

	return true
}