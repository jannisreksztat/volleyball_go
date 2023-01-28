package dvw

import (
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func CreateCSV(df dataframe.DataFrame, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	
	df.WriteCSV(f)
}
