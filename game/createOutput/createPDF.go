package createOutput

import (
	"math"
	"os"
	"path/filepath"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
)

func (output *Output) CreatePDF(directory string) {
	var pdfWidth float64 = 210
	var pdfHeight float64 = 297
	var scale float64 = 1

	directory = filepath.Join(directory, "volley.pdf")
	f, err := os.Create(directory)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p := pdf.New(f, pdfWidth, pdfHeight, nil)

	output.createField(pdfWidth, pdfHeight, scale)
	output.createArrows(pdfWidth, pdfHeight, scale)

	output.canvas.RenderTo(p)
	if err := p.Close(); err != nil {
		panic(err)
	}
}

func (output *Output) createField(pdfWidth float64, pdfHeight float64, scale float64) {
	feld := output.volley.VolleyField

	output.context.SetFillColor(canvas.White)
	output.context.SetStrokeColor(canvas.Black)
	output.context.DrawPath(0, 0, canvas.Rectangle(output.width, output.height))

	for i := 0; i < len(feld); i++ {
		output.context.DrawPath(feld[i].XPos*scale+pdfWidth/2,
			feld[i].YPos*scale+pdfHeight/2,
			canvas.Rectangle(feld[i].Size*scale,
				feld[i].Size*scale))
	}
}

func (output *Output) createArrows(pdfWidth float64, pdfHeight float64, scale float64) {
	arrows := output.volley.Aktion
	feld := output.volley.VolleyField

	var startpos = []float64{0, 0}
	var endpos = []float64{0, 0}

	for i := 0; i < len(arrows); i++ {
		for u := 0; u < len(feld); u++ {
			if arrows[i].Startzone == feld[u].Name {
				startpos[0] = feld[u].XPos + feld[u].Size/2
				startpos[1] = feld[u].YPos + feld[u].Size/2
			}

			if arrows[i].Endzone == feld[u].Name {
				endpos[0] = feld[u].XPos + feld[u].Size/2
				endpos[1] = feld[u].YPos + feld[u].Size/2
			}
		}
		output.drawArrow(startpos[0]*scale+pdfWidth/2, startpos[1]*scale+pdfHeight/2, endpos[0]*scale+pdfWidth/2, endpos[1]*scale+pdfHeight/2)
	}

}

func (output *Output) drawArrow(x1, y1, x2, y2 float64) {
	output.context.SetStrokeColor(canvas.Red)
	var angle float64 = math.Atan((y2 - y1) / (x2 - x1))

	if x2 < x1 {
		angle = angle + math.Pi
	}

	var newX1 float64 = x2 + 5*math.Cos(angle+degreeToRadian(150))
	var newY1 float64 = y2 + 5*math.Sin(angle+degreeToRadian(150))
	var newX2 float64 = x2 + 5*math.Cos(angle-degreeToRadian(150))
	var newY2 float64 = y2 + 5*math.Sin(angle-degreeToRadian(150))

	output.context.MoveTo(x1, y1)
	output.context.LineTo(x2, y2)
	output.context.LineTo(newX1, newY1)
	output.context.MoveTo(x2, y2)
	output.context.LineTo(newX2, newY2)
	output.context.Stroke()
}

func degreeToRadian(degree float64) float64 {
	return math.Pi / 180 * degree
}

func radianToDegree(radian float64) float64 {
	return radian * 180 / math.Pi
}
