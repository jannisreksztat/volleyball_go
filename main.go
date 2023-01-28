package main

import (
	"fmt"
	dvw "volleyball_go/read_dvw"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	generalData, gameData := dvw.ReadDVW("example_data/5004_Hin01_Karlsruhe_Delitzsch.dvw")
	fmt.Println(generalData[0])

	df := dataframe.LoadRecords(gameData)

	// fmt.Println(df)

	// Get column names
	// fmt.Println(df.Names())

	// Get column types
	// fmt.Println(df.Types())

	// Summary
	// fmt.Println(df.Describe())

	//Returns rows
	// df.Subset()
	df = dvw.DeleteEmpty(df, "undefined")
	dvw.CreateCSV(df, "volleyball.csv")

}
