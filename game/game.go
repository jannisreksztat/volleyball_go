package game

import (
	reader "volleyball_go/game/fileReader"

	"github.com/go-gota/gota/dataframe"
)

type Game struct {
	GameData     dataframe.DataFrame
	GeneralData  []string
	FilteredData dataframe.DataFrame
}

func NewGame(filename string) *Game {
	game := new(Game)
	dvw := reader.ReadDVWFile(filename)

	game.GameData = dvw.GameDataFrame
	game.GeneralData = dvw.GeneralData
	game.FilteredData = game.GameData

	return game
}
