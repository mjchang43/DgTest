package taskthree

import(
	"testing"
	"strings"
)

func TestSolution_1(t *testing.T) {
	s := Solution("ACCAABBC")
	
	if strings.Compare(s, "AC") != 0 {
		t.Error("Expected, ACCAABBC =>", s)
	}else {
		t.Log("taskthree.Solution PASS, ACCAABBC =>", s)
	}
}

func TestSolution_2(t *testing.T) {
	s := Solution("ABCBBCBA")
	
	if strings.Compare(s, "") != 0 {
		t.Error("Expected, ABCBBCBA =>", s)
	}else {
		t.Log("taskthree.Solution PASS, ABCBBCBA =>", s)
	}
}

func TestSolution_3(t *testing.T) {
	s := Solution("BABABA")
	
	if strings.Compare(s, "BABABA") != 0 {
		t.Error("Expected, BABABA =>", s)
	}else {
		t.Log("taskthree.Solution PASS, BABABA =>", s)
	}
}

