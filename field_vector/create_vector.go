package fieldVector

import (
	dataFunction "volleyball_go/data_functions"

	"github.com/go-gota/gota/dataframe"
)

func (volley *Volley) createVectors(df dataframe.DataFrame, attackTeam string) {
	for row := 0; row < df.Nrow(); row++ {
		var attack vector
		if dataFunction.GetData(df, "Start zone", row) == "" {
			attack.Startzone = "8"
		} else {
			attack.Startzone = dataFunction.GetData(df, "Start zone", row)
		}
		if dataFunction.GetData(df, "End zone", row) == "" {
			attack.Endzone = attackTeam + "8"
		} else {
			attack.Endzone = dataFunction.GetData(df, "Team", row) + dataFunction.GetData(df, "End zone", row) +
				dataFunction.GetData(df, "End zone +", row)
		}
		attack.Type = dataFunction.GetData(df, "Type", row)
		attack.Rating = dataFunction.GetData(df, "Rating", row)
		volley.Aktion = append(volley.Aktion, attack)
	}
}
