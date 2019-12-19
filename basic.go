package basicoperation

import "errors"

// Add returns the summation of two float64
func Add(a, b float64) float64 {
	return a + b
}

// Sub returns the subtraction of two float64
func Sub(a, b float64) float64 {
	return a - b
}

// Mul returns the multiplication of two float64
func Mul(a, b float64) float64 {
	return a * b
}

// Div returns the division of two float64
func Div(a, b float64) (float64, error) {
	if a == 0 && b == 0 {
		return 0, errors.New("ERROR: 0 / 0 is not defined")
	} else if b == 0 {
		return 0, errors.New("ERROR: float64 / 0 is Infinit")
	}
	return a / b, nil
}
