// Math is a babel service to do basic math. This is the implementation.

package impl

import (
	"errors"
	"github.com/babelrpc/lib-go/examples/babeldemo/math"
	"math/big"
)

type FractionServiceImpl int

type frac math.Fraction

func (f frac) check() error {
	if f.Numerator == nil {
		return errors.New("Numerator cannot be null")
	} else if f.Denominator == nil {
		return errors.New("Denominator cannot be null")
	} else if *f.Denominator <= 0 {
		return errors.New("Denominator cannot be negative or zero")
	} else {
		return nil
	}
}

func (f frac) toRat() *big.Rat {
	R := big.NewRat(int64(*f.Numerator), int64(*f.Denominator))
	return R
}

func (f *frac) fromRat(r *big.Rat) {
	*f.Numerator = int32(r.Num().Int64())
	*f.Denominator = int32(r.Denom().Int64())
}

func (x FractionServiceImpl) check(a, b *math.Fraction) error {
	if a == nil || b == nil {
		return errors.New("Input values cannot be null")
	}
	err := frac(*a).check()
	if err != nil {
		return errors.New("a: " + err.Error())
	}
	err = frac(*b).check()
	if err != nil {
		return errors.New("b: " + err.Error())
	}
	return nil
}

// Adds a and b
// a:  first addend
// b:  second addend
func (x FractionServiceImpl) Add(a *math.Fraction, b *math.Fraction) (*math.Fraction, error) {
	err := x.check(a, b)
	if err != nil {
		return nil, err
	}
	result := new(math.Fraction)
	result.Init()
	A := frac(*a).toRat()
	B := frac(*b).toRat()
	C := frac(*result).toRat()
	C.Add(A, B)
	(*frac)(result).fromRat(C)
	return result, nil
}

// Subtracts b from a
// a:  minuend
// b:  subtrahend
func (x FractionServiceImpl) Subtract(a *math.Fraction, b *math.Fraction) (*math.Fraction, error) {
	err := x.check(a, b)
	if err != nil {
		return nil, err
	}
	result := new(math.Fraction)
	result.Init()
	A := frac(*a).toRat()
	B := frac(*b).toRat()
	C := frac(*result).toRat()
	C.Sub(A, B)
	(*frac)(result).fromRat(C)
	return result, nil
}

// Multiplies a and b
// a:  first factor
// b:  second factor
func (x FractionServiceImpl) Multiply(a *math.Fraction, b *math.Fraction) (*math.Fraction, error) {
	err := x.check(a, b)
	if err != nil {
		return nil, err
	}
	result := new(math.Fraction)
	result.Init()
	A := frac(*a).toRat()
	B := frac(*b).toRat()
	C := frac(*result).toRat()
	C.Mul(A, B)
	(*frac)(result).fromRat(C)
	return result, nil
}

// Divides a by b
// a:  dividend
// b:  divisor
func (x FractionServiceImpl) Divide(a *math.Fraction, b *math.Fraction) (result *math.Fraction, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	err = x.check(a, b)
	if err != nil {
		return nil, err
	}
	result = new(math.Fraction)
	result.Init()
	A := frac(*a).toRat()
	B := frac(*b).toRat()
	C := frac(*result).toRat()
	C.Quo(A, B)
	(*frac)(result).fromRat(C)
	return result, nil
}
