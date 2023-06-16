package createOutput

import (
	"os"
	"path/filepath"

	"github.com/tdewolff/canvas/renderers/pdf"
)

func (output *Output) CreatePDF(directory string) {
	var pdfWidth float64 = 210			//struct an zentraler stelle erstellen (verschiedene Maße?)
	var pdfHeight float64 = 297
	var scale float64 = 1

	directory = filepath.Join(directory, "volley.pdf")
	f, err := os.Create(directory)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	p := pdf.New(f, pdfWidth, pdfHeight, nil)

	output.createOutputData(pdfWidth, pdfHeight, scale)

	output.canvas.RenderTo(p)
	if err := p.Close(); err != nil {
		panic(err)
	}
}
