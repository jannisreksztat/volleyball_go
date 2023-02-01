package game

import (
	dataFunction "volleyball_go/data_functions"
	latexDraw "volleyball_go/latex_Draw"
)

func (game *Game) DrawMatch(size float32, attackTeam string, direction bool, directory string) {
	latexDraw.DrawLatex(size, attackTeam, direction, directory, game.createVectors())
}

func (game *Game) createVectors() []latexDraw.Vector {
	var vectors []latexDraw.Vector
	for row := 0; row < game.FilteredData.Nrow(); row++ {
		var attack latexDraw.Vector
		attack.Startzone = dataFunction.GetData(game.FilteredData, "Start zone", row)
		attack.Endzone = dataFunction.GetData(game.FilteredData, "End zone", row) + dataFunction.GetData(game.FilteredData, "End zone +", row)
		attack.Type = dataFunction.GetData(game.FilteredData, "Type", row)
		attack.Rating = dataFunction.GetData(game.FilteredData, "Rating", row)
		vectors = append(vectors, attack)
	}
	return vectors
}
