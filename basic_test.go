package basicoperation

import "testing"

// TestAdd tests the function Add
func TestAdd(t *testing.T) {
	addAnswer := Add(10, 20)
	if addAnswer != 30 {
		t.Errorf(" [ Error ] : Expected 30, Actual %d", addAnswer)
	}
}

// TestSub tests the function Sub
func TestSub(t *testing.T) {
	subAnswer := Sub(20, 20)
	if subAnswer != 0 {
		t.Errorf(" [ Error ] : Expected 0, Actual %d", subAnswer)
	}
}

// TestMul tests the function Mul
func TestMul(t *testing.T) {
	mulAnswer := Mul(20, 20)
	if mulAnswer != 400 {
		t.Errorf(" [ Error ] : Expected 400, Actual %d", mulAnswer)
	}
}

// TestDiv tests the function Div
func TestDiv(t *testing.T) {
	divAnswer, _ := Div(20, 20)
	if divAnswer != 1 {
		t.Errorf(" [ Error ] : Expected 1, Actual %d", divAnswer)
	}

	divA, actualErr := Div(10, 0)
	expectedErr := "ERROR: integer / 0 is Infinit"
	if divA != 0 || actualErr.Error() != expectedErr {
		t.Errorf("Expected: %s\nActual Error Message: %s", expectedErr, actualErr)
	}

}
