package fieldVector

import (
	dataFunction "volleyball_go/data_functions"

	"github.com/go-gota/gota/dataframe"
)

func (volley *Volley) createVectors(df dataframe.DataFrame, attackTeam string) {
	for row := 0; row < df.Nrow(); row++ {
		volley.Aktion = append(volley.Aktion, vector{
			volley.writeStartzone(row, df),
			volley.writeEndzone(row, df, attackTeam),
			volley.writeType(row, df),
			volley.writeRating(row, df),
		})
	}
}

func (volley *Volley) writeStartzone(row int, df dataframe.DataFrame) string {
	if dataFunction.GetData(df, "Start zone", row) == "" {
		return "8"
	} else {
		return dataFunction.GetData(df, "Start zone", row)
	}
}

func (volley *Volley) writeEndzone(row int, df dataframe.DataFrame, attackTeam string) string {
	if dataFunction.GetData(df, "End zone", row) == "" {
		return attackTeam + "8"
	} else {
		return dataFunction.GetData(df, "Team", row) + dataFunction.GetData(df, "End zone", row) +
			dataFunction.GetData(df, "End zone +", row)
	}
}

func (volley *Volley) writeType(row int, df dataframe.DataFrame) string {
	return dataFunction.GetData(df, "Type", row)
}

func (volley *Volley) writeRating(row int, df dataframe.DataFrame) string {
	rating := dataFunction.GetData(df, "Rating", row)
	if rating == "#" {
		return `\#`
	} else {
		return rating
	}

}
