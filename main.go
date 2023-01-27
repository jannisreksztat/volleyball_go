package main

import (
	"fmt"
	dvw "volleyball_go/read_dvw"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	generalData, gameData := dvw.ReadDVW("example_data/5004_Hin01_Karlsruhe_Delitzsch.dvw")
	// generalData, gameData := dvw.ReadDVW("example_data/aim.txt")
	fmt.Println(generalData[0])

	df := dataframe.LoadRecords(gameData)
	fmt.Println(df)

}
