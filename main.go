package main

import (
	dataFunction "volleyball_go/data_functions"
	"volleyball_go/game"
	"volleyball_go/game/createOutput"
)

func main() {
	outputDir := dataFunction.CreateDirectory()

	data := game.NewGame("example_data/example_data.dvw")
	data.FilterTeamPlayer("a", "11")
	data.FilterSkill("A")
	// fmt.Println(data.FilteredData)
	// data.CreateCSV(outputDir)
	data.DrawMatch(8, "a", outputDir)
	output := createOutput.NewOutput(data.FilteredData, 100, "a")
	output.CreatePDF(outputDir)
}
