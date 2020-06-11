package numberline

import (
	"fmt"
	"testing"
)

var rangeTest Range

func TestNewRange(t *testing.T) {

	checkRange := func(t *testing.T, got, want Range) {
		t.Helper()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	assertNoError := func(t *testing.T, err error) {
		t.Helper()

		if err != nil {
			t.Errorf("got an error, but didn't want one")
		}

	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("didn't get an error, but wanted one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	var positiveTests = []struct {
		input string
		err   error
		want  Range
	}{
		{input: "(2,5]", err: nil, want: Range{LowerBound: 3, UpperBound: 5}},
		{input: "[-2,15)", err: nil, want: Range{LowerBound: -2, UpperBound: 14}},
	}

	for _, test := range positiveTests {
		message := fmt.Sprintf("valid range %v", test.want)
		t.Run(message, func(t *testing.T) {
			got, err := rangeTest.NewRange(test.input)
			checkRange(t, got, test.want)
			assertNoError(t, err)
		})
	}

	var negativeTests = []struct {
		input string
		err   error
		want  Range
	}{
		{input: "-2,15", err: errInvalidRange, want: Range{0, 0}},
		{input: "(-f,15)", err: errInvalidValue, want: Range{0, 0}},
	}

	for _, test := range negativeTests {
		message := fmt.Sprintf("invalid range %q", test.input)
		t.Run(message, func(t *testing.T) {
			got, err := rangeTest.NewRange(test.input)
			checkRange(t, got, Range{})
			assertError(t, err, test.err)
		})
	}
}

func TestContains(t *testing.T) {

	assertContains := func(t *testing.T, got, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	var tests = []struct {
		input     []int
		want      bool
		rangeTest Range
	}{
		{input: []int{3, 2, 5, 7}, want: false, rangeTest: Range{3, 5}},
		{input: []int{3, 3, 5, 4, 5}, want: true, rangeTest: Range{3, 5}},
		{input: []int{3, 5, 6}, want: false, rangeTest: Range{3, 5}},
		{input: []int{4}, want: true, rangeTest: Range{3, 5}},
	}

	for _, test := range tests {
		message := fmt.Sprintf("Numbers: %v", test.input)
		t.Run(message, func(t *testing.T) {
			got := test.rangeTest.Contains(test.input...)
			assertContains(t, got, test.want)
		})
	}

}

func TestDoesNotContain(t *testing.T) {
	rangeTest = Range{3, 5}

	assertNotContain := func(t *testing.T, got, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	var tests = []struct {
		input     []int
		want      bool
		rangeTest Range
	}{
		{input: []int{-1, 2, 7, 5}, want: true, rangeTest: Range{3, 5}},
		{input: []int{3, 5}, want: false, rangeTest: Range{3, 5}},
		{input: []int{1, 6}, want: true, rangeTest: Range{3, 5}},
		{input: []int{5}, want: false, rangeTest: Range{3, 5}},
	}

	for _, test := range tests {
		message := fmt.Sprintf("Numbers: %v", test.input)
		t.Run(message, func(t *testing.T) {
			got := test.rangeTest.DoesNotContain(test.input...)
			assertNotContain(t, got, test.want)
		})
	}

}

func TestGetAllPoints(t *testing.T) {

	assertPoints := func(t *testing.T, got, want []int) {
		t.Helper()

		if !compare(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	}

	var tests = []struct {
		rangeTest Range
		want      []int
	}{
		{rangeTest: Range{3, 5}, want: []int{3, 4, 5}},
		{rangeTest: Range{0, 7}, want: []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{rangeTest: Range{5, 5}, want: []int{5}},
		{rangeTest: Range{0, 0}, want: []int{0}},
	}

	for _, test := range tests {
		message := fmt.Sprintf("Range %v", test.rangeTest)
		t.Run(message, func(t *testing.T) {
			got := test.rangeTest.GetAllPoints()
			assertPoints(t, got, test.want)
		})
	}

}

func TestContainsRange(t *testing.T) {

	assertContainsRange := func(t *testing.T, got, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	}

	var tests = []struct {
		input     string
		rangeTest Range
		want      bool
	}{
		{input: "(2,6]", rangeTest: Range{3, 5}, want: false},
		{input: "(2,4]", rangeTest: Range{3, 5}, want: true},
		{input: "[0,5)", rangeTest: Range{1, 10}, want: false},
		{input: "[2,10]", rangeTest: Range{1, 10}, want: true},
	}

	for _, test := range tests {
		message := fmt.Sprintf("Comparing Range: %q", test.input)
		t.Run(message, func(t *testing.T) {
			got := test.rangeTest.ContainsRange(test.input)
			assertContainsRange(t, got, test.want)
		})
	}

}

func TestDoesNotContainRange(t *testing.T) {

	assertNotContainRange := func(t *testing.T, got, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	var tests = []struct {
		input     string
		rangeTest Range
		want      bool
	}{
		{input: "(2,6]", rangeTest: Range{3, 5}, want: true},
		{input: "(2,4]", rangeTest: Range{3, 5}, want: false},
		{input: "(0,5)", rangeTest: Range{1, 10}, want: false},
		{input: "(-1,7)", rangeTest: Range{1, 10}, want: true},
	}

	for _, test := range tests {
		message := fmt.Sprintf("Comparing range: %q", test.input)
		t.Run(message, func(t *testing.T) {
			got := test.rangeTest.DoesNotContainRange(test.input)
			assertNotContainRange(t, got, test.want)
		})
	}

}

func TestGetEndPoints(t *testing.T) {

	assertEndPoints := func(t *testing.T, lower, upper int, want Range) {
		t.Helper()

		if lower != want.LowerBound && upper != want.UpperBound {
			t.Errorf("got (%v,%v), want (%v,%v)", lower, upper, want.LowerBound, want.UpperBound)
		}
	}

	var tests = []struct {
		rangeTest Range
		lower     int
		upper     int
	}{
		{rangeTest: Range{3, 5}, lower: 3, upper: 5},
		{rangeTest: Range{1, 7}, lower: 1, upper: 7},
	}

	for _, test := range tests {
		message := fmt.Sprintf("Range %v", test.rangeTest)
		t.Run(message, func(t *testing.T) {
			lower, upper := test.rangeTest.GetEndPoints()
			assertEndPoints(t, lower, upper, test.rangeTest)
		})
	}

}

func TestOverlapsRange(t *testing.T) {

	assertOverlap := func(t *testing.T, got, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	var tests = []struct {
		input     string
		rangeTest Range
		want      bool
	}{
		{input: "(4,9]", rangeTest: Range{3, 5}, want: true},
		{input: "(5,8)", rangeTest: Range{3, 5}, want: false},
		{input: "[-1,3]", rangeTest: Range{3, 5}, want: true},
		{input: "(-2,2]", rangeTest: Range{3, 5}, want: false},
	}

	for _, test := range tests {
		message := fmt.Sprintf("Comparing range: %v", test.input)
		t.Run(message, func(t *testing.T) {
			got := test.rangeTest.OverlapsRange(test.input)
			assertOverlap(t, got, test.want)
		})
	}

}

func TestEquals(t *testing.T) {

	assertEquals := func(t *testing.T, got, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	var tests = []struct {
		input     string
		rangeTest Range
		want      bool
	}{
		{input: "(2,5]", rangeTest: Range{3, 5}, want: true},
		{input: "(3,5]", rangeTest: Range{3, 5}, want: false},
		{input: "[3,6)", rangeTest: Range{3, 5}, want: true},
		{input: "(-1,3]", rangeTest: Range{3, 5}, want: false},
	}

	for _, test := range tests {
		message := fmt.Sprintf("Comparing range: %v", test.input)
		t.Run(message, func(t *testing.T) {
			got := test.rangeTest.Equals(test.input)
			assertEquals(t, got, test.want)
		})
	}
}

func TestNotEquals(t *testing.T) {
	rangeTest = Range{3, 5}

	assertNotEquals := func(t *testing.T, got, want bool) {
		t.Helper()

		if got != want {
			t.Errorf("got %v, want %v", want, got)
		}
	}

	var tests = []struct {
		input     string
		rangeTest Range
		want      bool
	}{
		{input: "(2,5]", rangeTest: Range{3, 5}, want: false},
		{input: "(3,5]", rangeTest: Range{3, 5}, want: true},
		{input: "(-1,3]", rangeTest: Range{3, 5}, want: true},
		{input: "[3,6)", rangeTest: Range{3, 5}, want: false},
	}

	for _, test := range tests {
		message := fmt.Sprintf("Comparing range: %v", test.input)
		t.Run(message, func(t *testing.T) {
			got := test.rangeTest.NotEquals(test.input)
			assertNotEquals(t, got, test.want)
		})
	}

}

func compare(got, want []int) bool {

	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}

	return true

}
