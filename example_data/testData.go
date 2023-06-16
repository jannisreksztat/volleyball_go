package exampleData

import (
	"fmt"

	"github.com/go-gota/gota/dataframe"
)

func ExampleDF() dataframe.DataFrame {
	return dataframe.LoadRecords(
		[][]string{
			[]string{"A", "B", "C", "D"},
			[]string{"a", "4", "5.1", "true"},
			[]string{"k", "5", "7.0", "true"},
			[]string{"k", "4", "6.0", "true"},
			[]string{"a", "2", "7.1", "false"},
		},
	)
}

func ExampleVolleyDF() dataframe.DataFrame {
	return dataframe.LoadRecords(
		[][]string{
			[]string{
				"Team", "Player", "Skill", "Type", "Rating", "Cmb", "Targ Attack",
				"Start zone", "End zone", "End zone +", "Skill type+", "Players+", "Special", "undefined",
				"undefined", "undefined", "undefined", "undefined", "undefined", "Real Time", "Set",
				"P Feeder Home", "P Feeder Away", "VideoNr", "Timestamp Video", "undefined", "Home 1", "Home 2",
				"Home 3", "Home 4", "Home 5", "Home 6", "Away 1",
				"Away 2", "Away 3", "Away 4", "Away 5", "Away 6",
			},
			[]string{"a", "11", "R", "M", "-", "  ", " ", "1", "9", " ", " ", " ", " ", "", "", "", "", "", "", "14.31.34", "1", "3", "3", "", "", "", "6", "15", "4", "12", "10", "9",
				"11", "15", "10", "13", "3", "16"},
			[]string{"a", "11", "D", "H", "#", "  ", " ", "4", "6", " ", " ", " ", " ", "", "", "", "", "", "", "14.31.50", "1", "3", "3", "", "", "", "6", "15", "4", "12", "10", "9",
				"11", "15", "10", "13", "3", "16"},
		},
	)
}

func ExampleAttackString() []string {
	return []string{"*16RT#~~~68D", "*12SM#~~~95", "*09AH+VI~21~H"}
}

func CheckString(string1 string, string2 string) bool {
	if len(string1) != len(string2) {
		return false
	}
	for pos, con := range string1 {
		if con != rune(string2[pos]) {
			fmt.Printf("String stimmt nicht an Position: %d. Inhalt gewollt: %d, Wirklicher Inhalt: %d\n", pos, con, string2[pos])
			return false
		}
	}
	return true
}

func CheckStringArray(array1 []string, array2 []string) bool {
	if len(array1) != len(array2) {
		fmt.Printf("Array Laenge stimmt nicht! Erwartete laenge: %d, Wirkliche Laenge: %d\n", len(array1), len(array2))
		return false
	}
	for pos := 0; pos < len(array1); pos++ {
		if !CheckString(array1[pos], array2[pos]) {
			fmt.Printf("String stimmt nicht an Position: %d\n", pos)
			return false
		}
	}
	return true
}
