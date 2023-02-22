package createOutput

import (
	fieldVector "volleyball_go/game/createOutput/field_vector"

	"github.com/go-gota/gota/dataframe"
	"github.com/tdewolff/canvas"
)

type Output struct {
	volley  *fieldVector.Volley
	canvas  *canvas.Canvas
	context *canvas.Context
	width   float64
	height  float64
}

func NewOutput(gameData dataframe.DataFrame, size float64, attackTeam string) *Output {
	output := new(Output)

	output.volley = fieldVector.NewVolley(gameData, size, attackTeam)

	output.canvas = canvas.New(output.width, output.height)
	output.context = canvas.NewContext(output.canvas)

	output.width = 210
	output.height = 296

	return output
}
