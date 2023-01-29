package game

import (
	dvw "volleyball_go/read_dvw"

	"github.com/go-gota/gota/dataframe"
)

type Game struct {
	GeneralData  []string
	GameData     dataframe.DataFrame
	FilteredData dataframe.DataFrame
}

func NewGame(filename string) *Game {
	game := new(Game)
	var data [][]string
	game.GeneralData, data = dvw.ReadDVW(filename)
	game.GameData = dataframe.LoadRecords(data)
	game.FilteredData = game.GameData
	return game
}

// func (game Game) FilterPlayer(value string) {
// 	game.FilteredData = game.FilteredData.Filter(
// 		dataframe.F{
// 			Colname:    "Player",
// 			Comparator: series.Eq,
// 			Comparando: value,
// 		},
// 	)
// }
