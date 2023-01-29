package main

import (
	"fmt"
	"volleyball_go/game"
	dvw "volleyball_go/read_dvw"
)

func main() {
	// spiel := game.NewGame("example_data/5004_Hin01_Karlsruhe_Delitzsch.dvw")
	data := dvw.NewDVW("example_data/5004_Hin01_Karlsruhe_Delitzsch.dvw")
	fmt.Println(data.GameDataFrame)
	game.CreateCSV(data.GameDataFrame, "volleyball.csv")
	// spiel.FilterPlayer("5")
	// df = dvw.DeleteEmpty(df, "undefined")
	// df = dvw.DeleteRows(df)
	// df = dvw.FilterTeam(df, "a")
}
