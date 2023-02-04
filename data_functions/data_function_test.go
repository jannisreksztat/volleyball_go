package dataFunction

import (
	"fmt"
	"testing"
	exampleData "volleyball_go/example_data"
)

func ExampleGetData() {
	fmt.Println(GetData(exampleData.ExampleDF(), "D", 1))

	//Output:
	//true
}

func TestGetData(t *testing.T) {
	df := exampleData.ExampleDF()

	want := "true"
	result := GetData(df, "D", 2)

	if result != want {
		t.Errorf("Falscher Output! Erwarteter Output : %s, Realer Output : %s", want, result)
	}
	if result != want {
		t.Errorf("Falscher Output! Erwarteter Output : False")
	}

}

func TestCheckString(t *testing.T) {
	if CheckString("Test123") {
		t.Errorf("Falscher Output! Erwarteter Output : False")
	}

	if !CheckString("123235345") {
		t.Errorf("Falscher Output! Erwarteter Output : True")
	}

}

func TestCheckRune(t *testing.T) {
	if CheckRune('A') {
		t.Errorf("Falscher Output! Erwarteter Output : False")
	}
	if !CheckRune('1') {
		t.Errorf("Falscher Output! Erwarteter Output : True")
	}
}
