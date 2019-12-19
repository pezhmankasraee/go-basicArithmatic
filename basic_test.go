package basicoperation

import "testing"

var mockOperator OperatorPair = OperatorPair{A: 20.0, B: 30.0}

// TestAdd tests the function Add
func TestAdd(t *testing.T) {
	addAnswer := mockOperator.Add()
	if addAnswer != 50 {
		t.Errorf(" [ Error ] : Expected 50, Actual %f", addAnswer)
	}
}

// TestSub tests the function Sub
func TestSub(t *testing.T) {

	subAnswer := mockOperator.Sub()
	if subAnswer != -10 {
		t.Errorf(" [ Error ] : Expected -10, Actual %f", subAnswer)
	}
}

// TestMul tests the function Mul
func TestMul(t *testing.T) {
	mulAnswer := mockOperator.Mul()
	if mulAnswer != 600 {
		t.Errorf(" [ Error ] : Expected 600, Actual %f", mulAnswer)
	}
}

// TestDiv tests the function Div
func TestDiv(t *testing.T) {
	mockOperator = OperatorPair{A: 10, B: 10}
	divAnswer, _ := mockOperator.Div()
	if divAnswer != 1 {
		t.Errorf(" [ Error ] : Expected 0.6, Actual %f", divAnswer)
	}

	mockOperator = OperatorPair{A: 10, B: 0.0}
	divA, actualErr := mockOperator.Div()
	expectedErr := "ERROR: float64 / 0 is Infinit"
	if divA != 0 || actualErr.Error() != expectedErr {
		t.Errorf("Expected: %s\nActual Error Message: %s", expectedErr, actualErr)
	}

	mockOperator = OperatorPair{A: 0.0, B: 0.0}
	divA, actualErr = mockOperator.Div()
	expectedErr = "ERROR: 0 / 0 is not defined"
	if divA != 0 || actualErr.Error() != expectedErr {
		t.Errorf("Expected: %s\nActual Error Message: %s", expectedErr, actualErr)
	}
}
