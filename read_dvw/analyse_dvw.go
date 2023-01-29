package dvw

import (
	"regexp"
	"strings"

	"github.com/go-gota/gota/dataframe"
)

// func analyse_dvw() {

// }

func DeleteEmpty(input dataframe.DataFrame, prefix string) dataframe.DataFrame {
	var keepStrings []string
	for _, col := range input.Names() {
		if !strings.HasPrefix(col, prefix) {
			keepStrings = append(keepStrings, col)
		}
	}
	return input.Select(keepStrings)
}

func DeleteRows(input dataframe.DataFrame) dataframe.DataFrame {
	var keepRows []int
	for row := 0; row < input.Nrow(); row++ {
		if CheckString(GetData(input, "Player", row)) {
			keepRows = append(keepRows, row)
		}
	}
	return input.Subset(keepRows)
}

func GetData(df dataframe.DataFrame, column string, row int) string {
	return df.Col(column).Elem(row).String()
}

func CheckString(input string) bool {
	for _, content := range input {
		if !CheckRune(content) {
			return false
		}
	}
	return true
}

func CheckRune(input rune) bool {
	return regexp.MustCompile(`\d`).MatchString(string(input))
}
