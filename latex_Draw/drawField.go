package latexDraw

type field struct {
	Size float32
	Name string
	XPos float32
	YPos float32
}

var mainField = [][]string{
	{"2", "3", "4"},
	{"9", "8", "7"},
	{"1", "6", "5"},
}

func createFields(size float32, direction bool, attackTeam string) []field {
	var fields []field
	size = size / 3

	for row, rowCon := range mainField {
		for column, colCon := range rowCon {
			if direction {
				name := attackTeam + colCon
				fields = append(fields,
					field{size,
						name,
						(size * (float32(column) - 1)),
						(size * (float32(row) + size/2)),
					})
				fields = append(fields, createSubfields(size, size*(float32(column)-1), size*float32(row)+size/2, name)...)
			} else {
				fields = append(fields,
					field{size,
						colCon,
						(size * (float32(column) - 1) * -1),
						(-size*float32(row) - size/2),
					})
			}
		}
	}
	return fields
}

var subField = [][]string{
	{"D", "A"},
	{"C", "B"},
}

func createSubfields(size float32, xpos float32, ypos float32, mainName string) []field {
	var subfields []field
	size = size / 2
	for subRow, subRowCon := range subField {
		for subCol, subColCon := range subRowCon {
			subfields = append(subfields,
				field{size,
					mainName + subColCon,
					(xpos - float32(subCol)*size + size/2),
					(ypos - float32(subRow)*size + size/2),
				})
		}

	}
	return subfields
}
