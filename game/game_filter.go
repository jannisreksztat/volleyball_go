package game

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func (game *Game) FilterPlayer(value string) {
	game.FilteredData = game.FilteredData.Filter(
		dataframe.F{
			Colname:    "Player",
			Comparator: series.Eq,
			Comparando: value,
		},
	)
}

func (game *Game) FilterTeamPlayer(team string, player string) {
	game.FilteredData = game.FilteredData.FilterAggregation(
		dataframe.And,
		dataframe.F{
			Colname:    "Team",
			Comparator: series.Eq,
			Comparando: team,
		},
		dataframe.F{
			Colname:    "Player",
			Comparator: series.Eq,
			Comparando: player,
		},
	)
}

func (game *Game) FilterSkill(skill string) {
	game.FilteredData = game.FilteredData.Filter(
		dataframe.F{
			Colname:    "Skill",
			Comparator: series.Eq,
			Comparando: skill,
		},
	)
}
