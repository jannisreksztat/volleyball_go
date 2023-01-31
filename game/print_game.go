package game

import (
	"log"
	"os"
	"path/filepath"
)

func (game *Game) CreateCSV(path string) {
	path = filepath.Join(path, "volleyball.csv")
	filename, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	game.FilteredData.WriteCSV(filename)
}
