package dvw

import (
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

// func DeleteRows(input dataframe.DataFrame) dataframe.DataFrame {
// 	var keepRows []int
// 	for row := 0; row < input.Nrow(); row++ {
// 		element := input.Elem(row, 1)
// 		if unicode.IsDigit(element) {
// 			keepRows = append(keepRows, row)
// 		}
// 	}
// 	return input.Subset(keepRows)
// }
