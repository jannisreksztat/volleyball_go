package fieldVector

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func (volley *Volley) CreateTex(directory string) {

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

	err = tmpl.Execute(f, volley)
	if err != nil {
		panic(err)
	}
}
