package main

import (
	"fmt"
	"volleyball_go/game"
)

func main() {
	data := game.NewGame("example_data/5004_Hin01_Karlsruhe_Delitzsch.dvw")
	data.FilterTeamPlayer("a", "15")
	data.FilterSkill("A")
	fmt.Println(data.FilteredData)
	data.CreateCSV()
}
