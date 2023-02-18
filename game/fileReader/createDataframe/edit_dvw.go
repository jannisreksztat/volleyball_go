package dvw

import (
	"strings"
	dataFunction "volleyball_go/data_functions"
)

func (dvw *DVW) deleteColumns(prefix string) {
	var keepStrings []string
	for _, col := range dvw.GameDataFrame.Names() {
		if !strings.HasPrefix(col, prefix) {
			keepStrings = append(keepStrings, col)
		}
	}
	dvw.GameDataFrame.Select(keepStrings)
}

func (dvw *DVW) deleteRows() {
	var keepRows []int
	for row := 0; row < dvw.GameDataFrame.Nrow(); row++ {
		if dataFunction.CheckString(dataFunction.GetData(dvw.GameDataFrame, "Player", row)) {
			keepRows = append(keepRows, row)
		}
	}
	dvw.GameDataFrame.Subset(keepRows)
}
