package game

import (
	"log"
	"os"
)

func (game Game) CreateCSV(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	game.GameData.WriteCSV(f)
}
