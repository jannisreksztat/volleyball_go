package game

import (
	"fmt"
	dvw "volleyball_go/read_dvw"

	"github.com/go-gota/gota/dataframe"
)

type Game struct {
	GameData     dataframe.DataFrame
	GeneralData  []string
	FilteredData dataframe.DataFrame
}

func NewGame(filename string) *Game {
	game := new(Game)
	dvw := dvw.NewDVW(filename)
	game.GameData = dvw.GameDataFrame
	game.GeneralData = dvw.GeneralData
	game.FilteredData = game.GameData
	fmt.Println(game.GameData)
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
