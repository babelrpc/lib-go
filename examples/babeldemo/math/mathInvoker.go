// Math is a babel service to do basic math.

package math

// *** AUTO-GENERATED FILE - DO NOT MODIFY ***
// *** Generated from math.babel ***

import (
)


// FloatServiceAddRequest is the request structure used for invoking the Add method on the FloatService service.
type FloatServiceAddRequest struct {

	// first addend 
	A *float64 `json:"a,omitempty"`

	// second addend 
	B *float64 `json:"b,omitempty"`

}

// Init sets default values for a AddRequest
func (obj *FloatServiceAddRequest) Init() *FloatServiceAddRequest {
	return obj
}

// FloatServiceAddResponse is the response structure used for invoking the Add method on the FloatService service.
type FloatServiceAddResponse struct {

	Value *float64 `json:"Value,omitempty"`

}

// Init sets default values for a AddResponse
func (obj *FloatServiceAddResponse) Init() *FloatServiceAddResponse {
	return obj
}


// FloatServiceSubtractRequest is the request structure used for invoking the Subtract method on the FloatService service.
type FloatServiceSubtractRequest struct {

	// minuend 
	A *float64 `json:"a,omitempty"`

	// subtrahend 
	B *float64 `json:"b,omitempty"`

}

// Init sets default values for a SubtractRequest
func (obj *FloatServiceSubtractRequest) Init() *FloatServiceSubtractRequest {
	return obj
}

// FloatServiceSubtractResponse is the response structure used for invoking the Subtract method on the FloatService service.
type FloatServiceSubtractResponse struct {

	Value *float64 `json:"Value,omitempty"`

}

// Init sets default values for a SubtractResponse
func (obj *FloatServiceSubtractResponse) Init() *FloatServiceSubtractResponse {
	return obj
}


// FloatServiceMultiplyRequest is the request structure used for invoking the Multiply method on the FloatService service.
type FloatServiceMultiplyRequest struct {

	// first factor 
	A *float64 `json:"a,omitempty"`

	// second factor 
	B *float64 `json:"b,omitempty"`

}

// Init sets default values for a MultiplyRequest
func (obj *FloatServiceMultiplyRequest) Init() *FloatServiceMultiplyRequest {
	return obj
}

// FloatServiceMultiplyResponse is the response structure used for invoking the Multiply method on the FloatService service.
type FloatServiceMultiplyResponse struct {

	Value *float64 `json:"Value,omitempty"`

}

// Init sets default values for a MultiplyResponse
func (obj *FloatServiceMultiplyResponse) Init() *FloatServiceMultiplyResponse {
	return obj
}


// FloatServiceDivideRequest is the request structure used for invoking the Divide method on the FloatService service.
type FloatServiceDivideRequest struct {

	// dividend 
	A *float64 `json:"a,omitempty"`

	// divisor 
	B *float64 `json:"b,omitempty"`

}

// Init sets default values for a DivideRequest
func (obj *FloatServiceDivideRequest) Init() *FloatServiceDivideRequest {
	return obj
}

// FloatServiceDivideResponse is the response structure used for invoking the Divide method on the FloatService service.
type FloatServiceDivideResponse struct {

	Value *float64 `json:"Value,omitempty"`

}

// Init sets default values for a DivideResponse
func (obj *FloatServiceDivideResponse) Init() *FloatServiceDivideResponse {
	return obj
}




// FractionServiceAddRequest is the request structure used for invoking the Add method on the FractionService service.
type FractionServiceAddRequest struct {

	// first addend 
	A *Fraction `json:"a,omitempty"`

	// second addend 
	B *Fraction `json:"b,omitempty"`

}

// Init sets default values for a AddRequest
func (obj *FractionServiceAddRequest) Init() *FractionServiceAddRequest {
	return obj
}

// FractionServiceAddResponse is the response structure used for invoking the Add method on the FractionService service.
type FractionServiceAddResponse struct {

	Value *Fraction `json:"Value,omitempty"`

}

// Init sets default values for a AddResponse
func (obj *FractionServiceAddResponse) Init() *FractionServiceAddResponse {
	return obj
}


// FractionServiceSubtractRequest is the request structure used for invoking the Subtract method on the FractionService service.
type FractionServiceSubtractRequest struct {

	// minuend 
	A *Fraction `json:"a,omitempty"`

	// subtrahend 
	B *Fraction `json:"b,omitempty"`

}

// Init sets default values for a SubtractRequest
func (obj *FractionServiceSubtractRequest) Init() *FractionServiceSubtractRequest {
	return obj
}

// FractionServiceSubtractResponse is the response structure used for invoking the Subtract method on the FractionService service.
type FractionServiceSubtractResponse struct {

	Value *Fraction `json:"Value,omitempty"`

}

// Init sets default values for a SubtractResponse
func (obj *FractionServiceSubtractResponse) Init() *FractionServiceSubtractResponse {
	return obj
}


// FractionServiceMultiplyRequest is the request structure used for invoking the Multiply method on the FractionService service.
type FractionServiceMultiplyRequest struct {

	// first factor 
	A *Fraction `json:"a,omitempty"`

	// second factor 
	B *Fraction `json:"b,omitempty"`

}

// Init sets default values for a MultiplyRequest
func (obj *FractionServiceMultiplyRequest) Init() *FractionServiceMultiplyRequest {
	return obj
}

// FractionServiceMultiplyResponse is the response structure used for invoking the Multiply method on the FractionService service.
type FractionServiceMultiplyResponse struct {

	Value *Fraction `json:"Value,omitempty"`

}

// Init sets default values for a MultiplyResponse
func (obj *FractionServiceMultiplyResponse) Init() *FractionServiceMultiplyResponse {
	return obj
}


// FractionServiceDivideRequest is the request structure used for invoking the Divide method on the FractionService service.
type FractionServiceDivideRequest struct {

	// dividend 
	A *Fraction `json:"a,omitempty"`

	// divisor 
	B *Fraction `json:"b,omitempty"`

}

// Init sets default values for a DivideRequest
func (obj *FractionServiceDivideRequest) Init() *FractionServiceDivideRequest {
	return obj
}

// FractionServiceDivideResponse is the response structure used for invoking the Divide method on the FractionService service.
type FractionServiceDivideResponse struct {

	Value *Fraction `json:"Value,omitempty"`

}

// Init sets default values for a DivideResponse
func (obj *FractionServiceDivideResponse) Init() *FractionServiceDivideResponse {
	return obj
}




// The Float service performs operations on floating-point values
type FloatService struct {
	SvcObj IFloatService `json:"-"`
}

// Adds a and b
func (s *FloatService) Add(req *FloatServiceAddRequest, rsp *FloatServiceAddResponse) error {
	response, err := s.SvcObj.Add(req.A, req.B)
	if err == nil {
		rsp.Value = response
	}
	return err
}

// Subtracts b from a
func (s *FloatService) Subtract(req *FloatServiceSubtractRequest, rsp *FloatServiceSubtractResponse) error {
	response, err := s.SvcObj.Subtract(req.A, req.B)
	if err == nil {
		rsp.Value = response
	}
	return err
}

// Multiplies a and b
func (s *FloatService) Multiply(req *FloatServiceMultiplyRequest, rsp *FloatServiceMultiplyResponse) error {
	response, err := s.SvcObj.Multiply(req.A, req.B)
	if err == nil {
		rsp.Value = response
	}
	return err
}

// Divides a by b
func (s *FloatService) Divide(req *FloatServiceDivideRequest, rsp *FloatServiceDivideResponse) error {
	response, err := s.SvcObj.Divide(req.A, req.B)
	if err == nil {
		rsp.Value = response
	}
	return err
}


// The Fraction service performs operations on fractional values
type FractionService struct {
	SvcObj IFractionService `json:"-"`
}

// Adds a and b
func (s *FractionService) Add(req *FractionServiceAddRequest, rsp *FractionServiceAddResponse) error {
	response, err := s.SvcObj.Add(req.A, req.B)
	if err == nil {
		rsp.Value = response
	}
	return err
}

// Subtracts b from a
func (s *FractionService) Subtract(req *FractionServiceSubtractRequest, rsp *FractionServiceSubtractResponse) error {
	response, err := s.SvcObj.Subtract(req.A, req.B)
	if err == nil {
		rsp.Value = response
	}
	return err
}

// Multiplies a and b
func (s *FractionService) Multiply(req *FractionServiceMultiplyRequest, rsp *FractionServiceMultiplyResponse) error {
	response, err := s.SvcObj.Multiply(req.A, req.B)
	if err == nil {
		rsp.Value = response
	}
	return err
}

// Divides a by b
func (s *FractionService) Divide(req *FractionServiceDivideRequest, rsp *FractionServiceDivideResponse) error {
	response, err := s.SvcObj.Divide(req.A, req.B)
	if err == nil {
		rsp.Value = response
	}
	return err
}



