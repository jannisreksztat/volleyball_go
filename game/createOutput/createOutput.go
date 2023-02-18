package createOutput

import (
	fieldVector "volleyball_go/game/createOutput/field_vector"

	"github.com/go-gota/gota/dataframe"
)

type Output struct {
	volley *fieldVector.Volley
}

func NewOutput(gameData dataframe.DataFrame, size float32, attackTeam string) *Output {
	output := new(Output)

	output.volley = fieldVector.NewVolley(gameData, size, attackTeam)

	return output
}
