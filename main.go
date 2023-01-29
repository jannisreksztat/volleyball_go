package main

import (
	"fmt"
	dvw "volleyball_go/read_dvw"
)

func main() {
	// spiel := game.NewGame("example_data/5004_Hin01_Karlsruhe_Delitzsch.dvw")
	dvw.NewDVW("example_data/5004_Hin01_Karlsruhe_Delitzsch.dvw")
	fmt.Printf(dvw.DVW.GameDataFrame)
	// spiel.CreateCSV("volleyball.csv")
	// spiel.FilterPlayer("5")
	// df = dvw.DeleteEmpty(df, "undefined")
	// df = dvw.DeleteRows(df)
	// df = dvw.FilterTeam(df, "a")
}
