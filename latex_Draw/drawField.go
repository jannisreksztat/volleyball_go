package latexDraw

import "fmt"

type field struct {
	Size float32
	Name string
	XPos float32
	YPos float32
}

var mainField = [][]int{
	{2, 3, 4},
	{9, 8, 7},
	{1, 6, 5},
}

func createFields(size float32, direction bool) []field {
	var fields []field
	for row, rowCon := range mainField {
		for column, colCon := range rowCon {
			name := fmt.Sprint(colCon)
			if direction {
				fields = append(fields,
					field{size / 3, name,
						(size / 3 * (float32(column) - 1)),
						(size/3*float32(row) + size/6)})
			} else {
				fields = append(fields,
					field{size / 3,
						name,
						(size * (float32(column) - 1) * -1 / 3),
						(-size*float32(row)/3 - size/6)})
			}
		}
	}
	return fields
}
