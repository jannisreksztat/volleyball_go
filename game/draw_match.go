package game

import (
	dataFunction "volleyball_go/data_functions"
	latexDraw "volleyball_go/latex_Draw"
)

func (game *Game) DrawMatch(size float32, attackTeam string, directory string) {
	latexDraw.DrawLatex(size, directory, game.createVectors(attackTeam), attackTeam)
}

func (game *Game) createVectors(attackTeam string) []latexDraw.Vector {
	var vectors []latexDraw.Vector
	for row := 0; row < game.FilteredData.Nrow(); row++ {
		var attack latexDraw.Vector
		if dataFunction.GetData(game.FilteredData, "Start zone", row) == "" {
			attack.Startzone = "8"
		} else {
			attack.Startzone = dataFunction.GetData(game.FilteredData, "Start zone", row)
		}
		if dataFunction.GetData(game.FilteredData, "End zone", row) == "" {
			attack.Endzone = attackTeam + "8"
		} else {
			attack.Endzone = dataFunction.GetData(game.FilteredData, "Team", row) + dataFunction.GetData(game.FilteredData, "End zone", row) + dataFunction.GetData(game.FilteredData,
				"End zone +", row)
		}
		attack.Type = dataFunction.GetData(game.FilteredData, "Type", row)
		attack.Rating = dataFunction.GetData(game.FilteredData, "Rating", row)
		vectors = append(vectors, attack)
	}
	return vectors
}
