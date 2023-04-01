package dvw

import (
	"fmt"
	"testing"
	testData "volleyball_go/example_data"
)

func ExampleSpecifyAttack() {
	fmt.Println(specifyAttack(testData.ExampleAttackString()[0]))
	//Output:
	//[* 16 R T #      6 8 D      ]
}

func TestSpecifyAttack(t *testing.T) {
	attackString := testData.ExampleAttackString()
	want := [][]string{
		{"*", "16", "R", "T", "#", "  ", " ", "6", "8", "D", " ", " ", " "},
		{"*", "12", "S", "M", "#", "  ", " ", "9", "5", " ", " ", " ", " "},
		{"*", "09", "A", "H", "+", "VI", " ", "2", "1", " ", "H", " ", " "},
	}
	var result []string

	for pos := 0; pos < len(attackString); pos++ {
		result = specifyAttack(attackString[pos])
		if !testData.CheckStringArray(want[pos], result) {
			t.Errorf("Arrays an Position: %d stimmen nicht ueberein\n", pos)
		}
	}
}

func TestSplitContent(t *testing.T) {
	want := []string{"*", "16", "R", "T", "#", "  ", " ", "6", "8", "D", " ", " ", " ", "1", "2", "3", "4"}
	result := splitContent("*16RT#~~~68D;1;2;3;4", ";")
	if !testData.CheckStringArray(want, result) {
		t.Errorf("Die Ergebnisse stimmen nicht ueberein")
	}
}
