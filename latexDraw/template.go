package latexDraw

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type volley struct {
	TikzTemplate string
	VolleyField  []field
}

type field struct {
	Size float32
	Name string
	XPos float32
	YPos float32
}

func DrawTemplate(size float32, attackTeam string, direction bool, directory string) {
	// feld := []field{
	// 	{size, "*", ypos},
	// 	{size, "a", -ypos},
	// }

	feld := createFields(size, attackTeam, direction)
	fmt.Println(feld)
	data := volley{
		tikz_template,
		feld,
	}

	tmpl, err := template.New("volley").Parse(volley_template)
	if err != nil {
		panic(err)
	}

	directory = filepath.Join(directory, "volley.tex")
	f, err := os.Create(directory)
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

// var names = []string{
// 	"A", "B", "C", "D",
// }

var pos = [][]int{
	{2, 3, 4},
	{9, 8, 7},
	{1, 6, 5},
}

func createFields(size float32, team string, direction bool) []field {
	var fields []field
	for row, rowCon := range pos {
		for column, colCon := range rowCon {
			// fmt.Println("Row: ", float32(row), "Column: ", float32(column))
			fmt.Println((size * (float32(column) - 1) * 1 / 3))
			name := team + fmt.Sprint(colCon)
			if direction {
				fields = append(fields, field{size / 3, name,
					(size * (float32(column) - 1) * 1 / 3), (size * float32(row) * 1 / 3)})
			} else {
				fields = append(fields, field{size / 3, name,
					1 / 3 * size * (float32(-column) - 1), 1 / 3 * size * float32(-row)})
			}
		}
	}
	return fields
}
