package numberline

import (
	"errors"
	"log"
	"strconv"
)

// Range data type
type Range struct {
	LowerBound, UpperBound int
}

var (
	errInvalidRange error = errors.New("Error: Invalid range limits")
	errInvalidValue error = errors.New("Error: Invalid value")
)

// NewRange creates a new instance of type Range
func (r Range) NewRange(expression string) (Range, error) {
	var numRange Range

	validLowerRange := [2]byte{40, 91}
	validUpperRange := [2]byte{41, 93}

	var isLowerRange bool = contains(validLowerRange, expression[0])
	var isUpperRange bool = contains(validUpperRange, expression[len(expression)-1])

	if !(isLowerRange && isUpperRange) {
		return numRange, errInvalidRange
	}

	var lowerBound, upperBound string
	var isLowerBound bool = true

	for i := 1; i < len(expression)-1; i++ {
		if expression[i] == 44 {
			isLowerBound = false
		} else if isLowerBound {
			lowerBound += string(expression[i])
		} else {
			upperBound += string(expression[i])
		}
	}

	lowerLimit, err := strconv.Atoi(lowerBound)
	if err != nil {
		return numRange, errInvalidValue
	}

	upperLimit, err := strconv.Atoi(upperBound)
	if err != nil {
		return numRange, errInvalidValue
	}

	if expression[0] == validLowerRange[0] {
		lowerLimit++
	}

	if expression[len(expression)-1] == validUpperRange[0] {
		upperLimit--
	}

	numRange = Range{lowerLimit, upperLimit}

	return numRange, nil
}

// Contains checks if Range r contains all of the specified numbers
func (r Range) Contains(numbers ...int) bool {

	for _, number := range numbers {
		if !(r.LowerBound <= number && r.UpperBound >= number) {
			return false
		}
	}

	return true
}

// DoesNotContain checks if Range r doesn't contain all of the specified numbers
func (r Range) DoesNotContain(numbers ...int) bool {

	if r.Contains(numbers...) {
		return false
	}

	return true

}

// GetAllPoints returns a slice with all the numbers inside Range r
func (r Range) GetAllPoints() []int {
	var points []int

	for i := r.LowerBound; i <= r.UpperBound; i++ {
		points = append(points, i)
	}

	return points

}

// ContainsRange checks if Range r contains the specified expression
func (r Range) ContainsRange(expression string) bool {
	var comparingRange Range
	var err error

	comparingRange, err = comparingRange.NewRange(expression)
	if err != nil {
		log.Fatalln(err)
	}

	var isInsideRange bool = r.LowerBound <= comparingRange.LowerBound && r.UpperBound >= comparingRange.UpperBound

	if isInsideRange {
		return true
	}

	return false

}

// DoesNotContainRange checks if Range r doesn’t contain the specified expression
func (r Range) DoesNotContainRange(expression string) bool {

	if r.ContainsRange(expression) {
		return false
	}

	return true

}

// GetEndPoints returns the lower and upper bound of Range r
func (r Range) GetEndPoints() (lower, upper int) {
	lower = r.LowerBound
	upper = r.UpperBound
	return
}

// OverlapsRange checks if the lower and upper bound of the specified expression overlaps with Range r
func (r Range) OverlapsRange(expression string) bool {
	var comparingRange Range
	var err error

	comparingRange, err = comparingRange.NewRange(expression)
	if err != nil {
		log.Fatalln(err)
	}

	var aboveLowerBound bool = r.LowerBound == comparingRange.LowerBound || r.LowerBound <= comparingRange.UpperBound
	var belowUpperBound bool = r.UpperBound == comparingRange.LowerBound || r.UpperBound >= comparingRange.UpperBound

	if aboveLowerBound && belowUpperBound {
		return true
	}

	return false

}

// Equals checks if Range r is equal to the specified expression
func (r Range) Equals(expression string) bool {
	var comparingRange Range
	var err error

	comparingRange, err = comparingRange.NewRange(expression)
	if err != nil {
		log.Fatalln(err)
	}

	if r == comparingRange {
		return true
	}

	return false

}

// NotEquals checks if Range r is not equal to the specified expression
func (r Range) NotEquals(expression string) bool {

	if r.Equals(expression) {
		return false
	}

	return true
}

func contains(byteArr [2]byte, byteChar byte) bool {
	for _, char := range byteArr {
		if byteChar == char {
			return true
		}
	}

	return false
}
