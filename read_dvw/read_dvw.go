package read_dvw

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadDVW(dataName string) (string, string) {
	var generalData string
	var gameData string

	data, err := os.Open(dataName)

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(data)

	scanner.Split(SplitGame)
	for scanner.Scan() {
		if generalData == "" {
			generalData = scanner.Text()
		} else {
			gameData = scanner.Text()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	gameData = strings.ReplaceAll(gameData, ";", ",")

	return generalData, gameData
}

func SplitGame(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the index of 3Scout to split
	if i := strings.Index(string(data), "[3SCOUT]"); i >= 0 {
		return i + 8, data[0:(i + 8)], nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), data, nil
	}
	return
}
