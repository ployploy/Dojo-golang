package test

import (
	"testing"
)

func TestInput1_whendivide3_expectedShownodisplay(t *testing.T) {
	result := divide(1)
	expected := "no display"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}

func TestInput3_whendivide3_expectedShow3(t *testing.T) {
	result := divide(3)
	expected := "3"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}

func TestInput102_whendivide3_expectedNotshow(t *testing.T) {
	result := divide(102)
	expected := "no show"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}
