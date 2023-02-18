package game

import (
	createOutput "volleyball_go/game/createOutput"
)

func (game *Game) DrawMatch(size float32, attackTeam string, directory string) {
	output := createOutput.NewOutput(game.FilteredData, size, attackTeam)
	output.CreateTex(directory)
}
