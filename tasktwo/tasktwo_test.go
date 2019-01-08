package tasktwo

import (
	"testing"
)

func TestSolution_1(t *testing.T) {
	A := []int{2, 8, 4, 3, 2}
	x := Solution(A, 7, 11, 3)

	if x != 8 {
		t.Error("Expected, got", A, 7, 11, 3, x)
	} else {
		t.Log("tasktwo.Solution PASS", A, 7, 11, 3, x)
	}
}

func TestSolution_2(t *testing.T) {
	A := []int{5}
	x := Solution(A, 4, 0, 3)
	
	if x != -1 {
		t.Error("Expected, got", A, 4, 0, 3, x)
	} else {
		t.Log("tasktwo.Solution PASS", A, 4, 0, 3, x)
	}
}

func TestSolution_3(t *testing.T) {
	A := []int{2, 8, 4}
	x := Solution(A, 7, 11, 3)
	
	if x != 2 {
		t.Error("Expected, got", A, 7, 11, 3, x)
	} else {
		t.Log("tasktwo.Solution PASS", A, 7, 11, 3, x)
	}
}
