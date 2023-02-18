package fileReader

import (
	"log"
	"os"
	dvw "volleyball_go/createDataframe"
)

func ReadDVWFile(filename string) *dvw.DVW {
	data, err := os.Open(filename) //file öffnen exportieren

	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	fileData := dvw.NewDVW(data)
	return fileData
}
