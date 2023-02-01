package latexDraw

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type volley struct {
	TikzTemplate string
	FeldLinie    float32
	VolleyField  []field
	Aktion       []arrow
}

type arrow struct {
	Startwert string
	Endwert   string
	Name      string
}
type field struct {
	Size float32
	Name string
	XPos float32
	YPos float32
}

func DrawTemplate(size float32, attackTeam string, direction bool, directory string) {

	feld := createFields(size, attackTeam, direction)
	feld = append(feld, createFields(size, "", !direction)...)

	arrows := []arrow{
		{"2", "*4", "#"},
	}

	data := volley{
		tikz_template,
		size,
		feld,
		arrows,
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

var pos = [][]int{
	{2, 3, 4},
	{9, 8, 7},
	{1, 6, 5},
}

func createFields(size float32, team string, direction bool) []field {
	var fields []field
	for row, rowCon := range pos {
		for column, colCon := range rowCon {
			name := team + fmt.Sprint(colCon)
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
