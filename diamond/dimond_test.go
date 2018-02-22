package test

import (
	"fmt"
	"testing"
)

func Test_print_dimondB(t *testing.T) {
	result := printDimond("B")
	fmt.Printf(result)
	expected := " A \nB B\n A \n"
	fmt.Printf(expected)
	if result != expected {
		t.Errorf("result is not equal %s\n", expected)
	}
}

func Test_print_dimondC(t *testing.T) {
	result := printDimond("C")
	expected := "  A  \n B B \nC   C\n B B \n  A  \n"
	if result != expected {
		t.Errorf("result is not equal %s", expected)
	}
}

func Test_print_dimondD(t *testing.T) {
	result := printDimond("D")
	fmt.Printf(result)
	expected := "   A   \n  B B  \n C   C \nD     D\n C   C \n  B B  \n   A   \n"
	fmt.Printf(expected)
	if result != expected {
		t.Errorf("result is not equal %s", expected)
	}
}
