package fieldVector

func (volley *Volley) createFields(attackTeam string) {
	size := volley.FeldSize / 3

	for row, rowCon := range mainField {
		for column, colCon := range rowCon {
			name := attackTeam + colCon
			xpos := (size * (float64(column) - 1.5))
			ypos := (size*(float64(row)) + size/2)
			volley.VolleyField = append(volley.VolleyField,
				field{size,
					name,
					xpos,
					ypos,
				})

			volley.createSubfields(xpos, ypos, name)
			volley.VolleyField = append(volley.VolleyField,
				field{size,
					colCon,
					(size * (float64(column) - 0.5) * -1),
					(-size*float64(row) - size/2),
				})

		}
	}
}

func (volley *Volley) createSubfields(xpos float64, ypos float64, mainName string) {
	size := volley.FeldSize / 6
	for subRow, subRowCon := range subField {
		for subCol, subColCon := range subRowCon {
			volley.VolleyField = append(volley.VolleyField,
				field{size,
					mainName + subColCon,
					(xpos - float64(subCol)*size + size),
					(ypos - float64(subRow)*size + size),
				})
		}

	}
}
