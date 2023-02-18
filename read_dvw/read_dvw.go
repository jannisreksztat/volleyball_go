package dvw

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func (dvw *DVW) readDVW() {
	var attackData []string

	data, err := os.Open(dvw.FileName) //file öffnen exportieren

	if err != nil {
		log.Fatal(err)
	}									
	defer data.Close()
	scanner := bufio.NewScanner(data)

	var DataType bool = false
	for scanner.Scan() {
		switch DataType {
		case false:
			dvw.GeneralData = append(dvw.GeneralData, scanner.Text())

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
	dvw.createMatrix(attackData)
	return
}

func (dvw *DVW) createMatrix(dataString []string) {
	dvw.gameData = append(dvw.gameData, dvw.header)
	for _, content := range dataString {
		dvw.gameData = append(dvw.gameData, splitContent(content, ";"))
	}
}

func splitContent(input string, delimiter string) []string {
	splitString := strings.Split(input, delimiter)
	output := specifyAttack(splitString[0])

	for _, content := range splitString[1:] {
		output = append(output, content)
	}
	return output
}

func specifyAttack(input string) []string {
	input = strings.ReplaceAll(input, "~", " ")
	var output []string

	oldPos := 0
	positionSplit := []int{1, 3, 4, 5, 6, 8, 9, 10, 11, 12, 13, 14, 15}
	for _, content := range positionSplit {
		if (content - 1) < len(input) {
			output = append(output, input[oldPos:content])
			oldPos = content
		} else {
			output = append(output, " ")
		}
	}
	return output
}
