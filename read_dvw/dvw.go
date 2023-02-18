package dvw

import "github.com/go-gota/gota/dataframe"

type DVW struct {
	GameDataFrame dataframe.DataFrame
	GeneralData   []string
	FileName      string

	header   []string
	gameData [][]string
}

//hier io.reader als übergabe
func NewDVW(filename string) *DVW {
	dvw := new(DVW)
	dvw.FileName = filename
	dvw.declareHeader()
	dvw.readDVW()
	dvw.GameDataFrame = dataframe.LoadRecords(dvw.gameData)

	dvw.deleteColumns("undefined")
	dvw.deleteRows()
	return dvw
}

func (dvw *DVW) declareHeader() {
	dvw.header = []string{
		"Team", "Player", "Skill", "Type", "Rating", "Cmb", "Targ Attack",
		"Start zone", "End zone", "End zone +", "Skill type+", "Players+", "Special", "undefined",
		"undefined", "undefined", "undefined", "undefined", "undefined", "Real Time", "Set",
		"P Feeder Home", "P Feeder Away", "VideoNr", "Timestamp Video", "undefined", "Home 1", "Home 2",
		"Home 3", "Home 4", "Home 5", "Home 6", "Away 1",
		"Away 2", "Away 3", "Away 4", "Away 5", "Away 6",
	}
	return
}
