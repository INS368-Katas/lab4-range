package numberline

import (
	"strconv"
	"errors"
	"log"
)

// Range data type
type Range struct {
	LowerBound, UpperBound int
}

var (
	errInvalidRange error = errors.New("error: invalid range limits")
	errInvalidValue error = errors.New("error: invalid value")
)


// NewRange creates a new instance of type Range
func (r Range) NewRange(expression string) (Range, error) {
	
	validLowerRange := [2]byte{40, 91}
	validUpperRange := [2]byte{41, 93}

	isLowerRange := contains(expression[0], validLowerRange)
	isUpperRange := contains(expression[len(expression) - 1], validUpperRange)

	if !(isLowerRange && isUpperRange) {
		return Range{}, errInvalidRange
	}

	var isLowerBound bool = true

	var lowerBound, upperBound string

	for i := 1; i < len(expression) - 1; i++ {
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
		return Range{}, errInvalidValue
	}

	upperLimit, err := strconv.Atoi(upperBound)
	if err != nil {
		return Range{}, errInvalidValue
	}

	if expression[0] == 40 {
		lowerLimit++
	}

	if expression[len(expression) - 1] == 41 {
		upperLimit--
	}

	return Range{lowerLimit, upperLimit}, nil


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

// GetAllPoints returns all the value inside Range r
func (r Range) GetAllPoints() []int {
	var points []int

	for i := r.LowerBound; i <= r.UpperBound; i++ {
		points = append(points, i)
	}

	return points

}

// ContainsRange checks if Range r contains the range in the specified expression 
func (r Range) ContainsRange(expression string) bool {
	var comparingRange Range
	var err error

	comparingRange, err = comparingRange.NewRange(expression)
	if err != nil {
		log.Fatalln(err)
	}

	var insideRange bool = r.LowerBound <= comparingRange.LowerBound && r.UpperBound >= comparingRange.UpperBound

	if insideRange {
		return true
	}

	return false

}

// DoesNotContainRange checks if Range r doesn't contain the range in the specified expression 
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

	var containsLowerBound bool = r.LowerBound <= comparingRange.LowerBound && r.UpperBound >= comparingRange.LowerBound
	var containsUpperBound bool = r.LowerBound <= comparingRange.UpperBound && r.UpperBound >= comparingRange.UpperBound

	if containsLowerBound || containsUpperBound {
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

// NotEquals checks if Range r is equal to the specified expression
func (r Range) NotEquals(expression string) bool {
	
	if r.Equals(expression) {
		return false
	}

	return true

}


func contains(char byte, byteArr [2]byte) bool {
	for i := range byteArr {
		if char == byteArr[i] {
			return true
		}
	}

	return false
}