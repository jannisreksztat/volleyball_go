package fieldVector

import (
	"github.com/go-gota/gota/dataframe"
)

type Volley struct {
	FeldSize    float32
	VolleyField []field
	Aktion      []vector
}
type field struct {
	Size float32
	Name string
	XPos float32
	YPos float32
}

type vector struct {
	Startzone string
	Endzone   string
	Type      string
	Rating    string
}

var mainField = [][]string{
	{"2", "3", "4"},
	{"9", "8", "7"},
	{"1", "6", "5"},
}

var subField = [][]string{
	{"D", "A"},
	{"C", "B"},
}

func NewVolley(gameData dataframe.DataFrame, size float32, attackTeam string) *Volley {
	volley := new(Volley)
	volley.FeldSize = size
	volley.createFields(attackTeam)

	volley.createVectors(gameData, attackTeam)

	return volley
}
