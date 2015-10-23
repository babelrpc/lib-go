// Math is a babel service to do basic math.

package math

// *** AUTO-GENERATED FILE - DO NOT MODIFY ***
// *** Generated from math.babel ***

import (
)




// Fraction represents a fractional value
type Fraction struct {
	// Numerator
	Numerator *int32 `json:"Numerator,omitempty"`

	// Denominator
	Denominator *int32 `json:"Denominator,omitempty"`

}

// Init sets default values for a Fraction
func (obj *Fraction) Init() *Fraction {
	obj.Numerator = new(int32)
	*obj.Numerator = 0
	obj.Denominator = new(int32)
	*obj.Denominator = 1
	return obj
}
