package basicoperation

import "errors"

// Add returns the summation of two integers
func Add(a, b int) int {
	return a + b
}

// Sub returns the subtraction of two integers
func Sub(a, b int) int {
	return a - b
}

// Mul returns the multiplication of two integers
func Mul(a, b int) int {
	return a * b
}

// Div returns the division of two integers
func Div(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, errors.New("ERROR: 0 / 0 is not defined")
	} else if b == 0 {
		return 0, errors.New("ERROR: integer / 0 is Infinit")
	}
	return a / b, nil
}
