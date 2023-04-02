package createOutput

// import (
// 	"os"
// 	"path/filepath"
// 	"syscall/js"

// 	"github.com/tdewolff/canvas/renderers/htmlcanvas"
// )

// func (output *Output) CreateHTMLCanvas(directory string) {
// 	var canvasWidth float64 = 200
// 	var canvasHeight float64 = 100
// 	var scale float64 = 1

// 	directory = filepath.Join(directory, "volley.html")
// 	f, err := os.Create(directory)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	cvs := js.Global().Get("document").Call("getElementById", "canvas")
// 	c := htmlcanvas.New(cvs, 200, 100, 5.0)
// 	output.createOutputData(canvasWidth, canvasHeight, scale)

// 	output.canvas.RenderTo(c)
// 	if err := c.Close(); err != nil {
// 		panic(err)
// 	}
// }
