package game

import (
	fieldVector "volleyball_go/field_vector"
)

func (game *Game) DrawMatch(size float32, attackTeam string, directory string) {
	volley := fieldVector.NewVolley(game.FilteredData, size, attackTeam)
	volley.CreateTex(directory)
}
