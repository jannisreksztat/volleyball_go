package createOutput

import "os"

func createPDF() {
	f, err := os.Create("document.pdf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
