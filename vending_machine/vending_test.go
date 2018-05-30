package vendingmachine

import (
	"reflect"
	"testing"
)

func Test_Buy_SD(t *testing.T) {
	insert := []string{"T", "F", "TW", "O"}
	choosedItem := "SD"
	actual := BuyItem(insert, choosedItem)
	expected := []string{"SD"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("your expected is %s but actual is %s", expected, actual)
	}
}

func Test_Return_Coin(t *testing.T) {
	insert := []string{"T", "T", "F"}
	choosedItem := "Coin Return"
	actual := BuyItem(insert, choosedItem)
	expected := []string{"T", "T", "F"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("your expected is %s but actual is %s", expected, actual)
	}
}

func Test_Buy_CC(t *testing.T) {
	insert := []string{"T", "T"}
	choosedItem := "CC"
	actual := BuyItem(insert, choosedItem)
	expected := []string{"F", "TW", "O"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("your expected is %s but actual is %s", expected, actual)
	}
}
