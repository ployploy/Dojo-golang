package string

import "testing"

func TestInputNumber50And2And1And9ShouldReturn95021(t *testing.T) {
	num := []int{50, 2, 1, 9}
	result := ConvertString(num)
	expected := "95021"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}

func TestInputNumber5And50And56ShouldReturn56550(t *testing.T) {
	num := []int{5, 50, 56}
	result := ConvertString(num)
	expected := "56550"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}

func TestInputNumber2And5And6And78ShouldReturn78256(t *testing.T) {
	num := []int{2, 5, 6, 78}
	result := ConvertString(num)
	expected := "78256"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}

func TestInputNumber2And5And6And789ShouldReturn78256(t *testing.T) {
	num := []int{2, 5, 6, 789}
	result := ConvertString(num)
	expected := "789256"
	if result != expected {
		t.Fatal("Expected ", expected, " but actual ", result)
	}
}
