package dvw

import (
	"strings"
	dataFunction "volleyball_go/data_functions"

	"github.com/go-gota/gota/dataframe"
)

func DeleteColumns(input dataframe.DataFrame, prefix string) dataframe.DataFrame {
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
		if dataFunction.CheckString(dataFunction.GetData(input, "Player", row)) {
			keepRows = append(keepRows, row)
		}
	}
	return input.Subset(keepRows)
}
