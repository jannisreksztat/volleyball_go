package main

import (
	game "volleyball_go/game"
	dvw "volleyball_go/read_dvw"
)

func main() {
	spiel := game.NewGame("example_data/5004_Hin01_Karlsruhe_Delitzsch.dvw")
	// spiel.FilterPlayer("5")
	dvw.CreateCSV(spiel.FilteredData, "volleyball.csv")
	// df = dvw.DeleteEmpty(df, "undefined")
	// df = dvw.DeleteRows(df)
	// df = dvw.FilterTeam(df, "a")

}
