// Math is a babel service to do basic math.

package math

// *** AUTO-GENERATED FILE - DO NOT MODIFY ***
// *** Generated from math.babel ***

import (
)


// The Float service performs operations on floating-point values

type IFloatService interface { 
	// Adds a and b
	// a:  first addend 
	// b:  second addend 
	Add(a *float64, b *float64) (*float64, error) 

	// Subtracts b from a
	// a:  minuend 
	// b:  subtrahend 
	Subtract(a *float64, b *float64) (*float64, error) 

	// Multiplies a and b
	// a:  first factor 
	// b:  second factor 
	Multiply(a *float64, b *float64) (*float64, error) 

	// Divides a by b
	// a:  dividend 
	// b:  divisor 
	Divide(a *float64, b *float64) (*float64, error) 
}

// The Fraction service performs operations on fractional values

type IFractionService interface { 
	// Adds a and b
	// a:  first addend 
	// b:  second addend 
	Add(a *Fraction, b *Fraction) (*Fraction, error) 

	// Subtracts b from a
	// a:  minuend 
	// b:  subtrahend 
	Subtract(a *Fraction, b *Fraction) (*Fraction, error) 

	// Multiplies a and b
	// a:  first factor 
	// b:  second factor 
	Multiply(a *Fraction, b *Fraction) (*Fraction, error) 

	// Divides a by b
	// a:  dividend 
	// b:  divisor 
	Divide(a *Fraction, b *Fraction) (*Fraction, error) 
}
