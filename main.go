package main

import (
	dataFunction "volleyball_go/data_functions"
	"volleyball_go/latexDraw"
)

func main() {
	outputDir := dataFunction.Test()

	// data := game.NewGame("example_data/5004_Hin01_Karlsruhe_Delitzsch.dvw")
	// data.FilterTeamPlayer("a", "15")
	// data.FilterSkill("A")
	// fmt.Println(data.FilteredData)
	// data.CreateCSV(outputDir)
	latexDraw.DrawTemplate(3, "*", false, outputDir)
}
