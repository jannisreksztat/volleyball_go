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
	Aktion       []Vector
}

type Vector struct {
	Startzone string
	Endzone   string
	Type      string
	Rating    string
}

func DrawLatex(size float32, directory string, vectors []Vector, attackTeam string) {

	feld := createFields(size, true, attackTeam)
	feld = append(feld, createFields(size, false, "")...)

	data := volley{
		tikz_template,
		size,
		feld,
		vectors,
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
