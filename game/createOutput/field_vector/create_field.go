package fieldVector

func (volley *Volley) createFields(attackTeam string) {
	size := volley.FeldSize / 3

	for row, rowCon := range mainField {
		for column, colCon := range rowCon {
			name := attackTeam + colCon
			volley.VolleyField = append(volley.VolleyField,
				field{size,
					name,
					(size * (float64(column) - 1)),
					(size*(float64(row)) + size/2),
				})

			volley.createSubfields(size*(float64(column)-1), size*float64(row)+size/2, name)
			volley.VolleyField = append(volley.VolleyField,
				field{size,
					colCon,
					(size * (float64(column) - 1) * -1),
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
