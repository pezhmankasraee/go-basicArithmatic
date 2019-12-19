package basicoperation

import "errors"

// OperatorPair keeps both operators
type OperatorPair struct {
	A float64
	B float64
}

// Add returns the summation of two float64
func (o OperatorPair) Add() float64 {
	return o.A + o.B
}

// Sub returns the subtraction of two float64
func (o OperatorPair) Sub() float64 {
	return o.A - o.B
}

// Mul returns the multiplication of two float64
func (o OperatorPair) Mul() float64 {
	return o.A * o.B
}

// Div returns the division of two float64
func (o OperatorPair) Div() (float64, error) {
	if o.A == 0 && o.B == 0 {
		return 0, errors.New("ERROR: 0 / 0 is not defined")
	} else if o.B == 0 {
		return 0, errors.New("ERROR: float64 / 0 is Infinit")
	}
	return o.A / o.B, nil
}
