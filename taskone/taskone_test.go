package taskone

import (
	"testing"
)

func TestSolution_1(t *testing.T) {
	A := []int{2,1,3,5,4}
	x := Solution(A)
	
	if x != 3 {
		t.Error("Expected, got", A, x)
	} else {
		t.Log("taskone.Solution PASS")
	}
}

func TestSolution_2(t *testing.T) {
	A := []int{2,3,4,1,5}
	x := Solution(A)
	
	if x != 2 {
		t.Error("Expected, got", A, x)
	} else {
		t.Log("taskone.Solution PASS")
	}
}

func TestSolution_3(t *testing.T) {
	A := []int{1,3,4,2,5}
	x := Solution(A)
	
	if x != 3 {
		t.Error("Expected, got", A, x)
	} else {
		t.Log("taskone.Solution PASS")
	}
}

