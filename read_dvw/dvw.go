package dvw

import "github.com/go-gota/gota/dataframe"

type DVW struct {
	GameDataFrame dataframe.DataFrame
	GeneralData   []string
	FileName      string

	header   []string
	gameData [][]string
}

func NewDVW(filename string) *DVW {
	dvw := new(DVW)
	dvw.FileName = filename
	dvw.DeclareHeader()
	dvw.ReadDVW()
	dvw.GameDataFrame = dataframe.LoadRecords(dvw.gameData)

	dvw.GameDataFrame = DeleteColumns(dvw.GameDataFrame, "undefined")
	dvw.GameDataFrame = DeleteRows(dvw.GameDataFrame)
	return dvw
}

func (dvw *DVW) DeclareHeader() {
	dvw.header = []string{
		"Team", "Player", "Skill", "Type", "Rating", "Cmb", "Targ Attack",
		"Start zone", "End zone", "End zone +", "Skill type+", "Players+", "Special", "undefined",
		"undefined", "undefined", "undefined", "undefined", "undefined", "Time", "undefined",
		"undefined", "undefined", "undefined", "undefined", "undefined", "Aufstellung 1", "Aufstellung 2",
		"Aufstellung 3", "Aufstellung 4", "Aufstellung 5", "Aufstellung 6", "Aufstellung 7",
		"Aufstellung 8", "Aufstellung 9", "Aufstellung 10", "Aufstellung 11", "Aufstellung 12",
	}
	return
}
