// Math is a babel service to do basic math. This is the implementation

package impl

import (
	"errors"
)

type FloatServiceImpl int

func check(a, b *float64) error {
	if a == nil || b == nil {
		return errors.New("Input values cannot be null")
	}
	return nil
}

// Adds a and b
// a:  first addend
// b:  second addend
func (x FloatServiceImpl) Add(a *float64, b *float64) (*float64, error) {
	err := check(a, b)
	if err != nil {
		return nil, err
	}
	result := new(float64)
	*result = (*a) + (*b)
	return result, nil
}

// Subtracts b from a
// a:  minuend
// b:  subtrahend
func (x FloatServiceImpl) Subtract(a *float64, b *float64) (*float64, error) {
	err := check(a, b)
	if err != nil {
		return nil, err
	}
	result := new(float64)
	*result = (*a) - (*b)
	return result, nil
}

// Multiplies a and b
// a:  first factor
// b:  second factor
func (x FloatServiceImpl) Multiply(a *float64, b *float64) (*float64, error) {
	err := check(a, b)
	if err != nil {
		return nil, err
	}
	result := new(float64)
	*result = (*a) * (*b)
	return result, nil
}

// Divides a by b
// a:  dividend
// b:  divisor
func (x FloatServiceImpl) Divide(a *float64, b *float64) (*float64, error) {
	err := check(a, b)
	if err != nil {
		return nil, err
	}
	if *b == 0.0 {
		return nil, errors.New("Divisor is zero")
	}
	result := new(float64)
	*result = (*a) / (*b)
	return result, nil
}
