package read_dvw

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadDVW(dataName string) ([]string, [][]string) {
	var generalData []string
	var attackData []string

	data, err := os.Open(dataName)

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(data)

	var DataType bool = false
	for scanner.Scan() {
		switch DataType {
		case false:
			generalData = append(generalData, scanner.Text())

		case true:
			attackData = append(attackData, scanner.Text())
		}
		if scanner.Text() == "[3SCOUT]" {
			DataType = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	return generalData, createMatrix(attackData)
}

func createMatrix(dataString []string) [][]string {
	var matrix [][]string

	for _, content := range dataString {
		matrix = append(matrix, strings.Split(content, ";"))
	}

	return matrix
}
