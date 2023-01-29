package game

import (
	"log"
	"os"
)

func (game *Game) CreateCSV() {
	filename, err := os.Create("volleyball.csv")
	if err != nil {
		log.Fatal(err)
	}
	game.FilteredData.WriteCSV(filename)
}
