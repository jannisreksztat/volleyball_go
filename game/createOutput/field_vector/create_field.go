package fieldVector

func (volley *Volley) createFields(attackTeam string) {
	size := volley.FeldSize / 3

	for row, rowCon := range mainField {
		for column, colCon := range rowCon {
			name := attackTeam + colCon
			volley.VolleyField = append(volley.VolleyField,
				field{size,
					name,
					(size * (float32(column) - 1)),
					(size*(float32(row)) + size/2),
				})

			volley.createSubfields(size*(float32(column)-1), size*float32(row)+size/2, name)
			volley.VolleyField = append(volley.VolleyField,
				field{size,
					colCon,
					(size * (float32(column) - 1) * -1),
					(-size*float32(row) - size/2),
				})

		}
	}
}

func (volley *Volley) createSubfields(xpos float32, ypos float32, mainName string) {
	size := volley.FeldSize / 6
	for subRow, subRowCon := range subField {
		for subCol, subColCon := range subRowCon {
			volley.VolleyField = append(volley.VolleyField,
				field{size,
					mainName + subColCon,
					(xpos - float32(subCol)*size + size/2),
					(ypos - float32(subRow)*size + size/2),
				})
		}

	}
}
