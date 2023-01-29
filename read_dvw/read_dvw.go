package dvw

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
	return generalData, CreateMatrix(attackData)
}

func CreateMatrix(dataString []string) [][]string {
	var matrix [][]string
	header := []string{"Team", "Player", "Skill", "Type", "Rating", "Cmb", "Targ Attack",
		"Start zone", "End zone", "End zone +", "Skill type+", "Players+", "Special", "undefined",
		"undefined", "undefined", "undefined", "undefined", "undefined", "Time", "undefined",
		"undefined", "undefined", "undefined", "undefined", "undefined", "Aufstellung 1", "Aufstellung 2",
		"Aufstellung 3", "Aufstellung 4", "Aufstellung 5", "Aufstellung 6", "Aufstellung 7",
		"Aufstellung 8", "Aufstellung 9", "Aufstellung 10", "Aufstellung 11", "Aufstellung 12",
	}

	matrix = append(matrix, header)
	for _, content := range dataString {
		matrix = append(matrix, SplitContent(content, ";"))
	}

	return matrix
}

func SplitContent(input string, delimiter string) []string {
	splitString := strings.Split(input, ";")
	output := SpecifyAttack(splitString[0])

	for _, content := range splitString[1:] {
		output = append(output, content)
	}
	return output
}

func SpecifyAttack(input string) []string {
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
