package main

import (
	"fmt"
	dvw "volleyball_go/read_dvw"
)

func main() {
	var generalData, gameData string = dvw.ReadDVW("example_data/aim.txt")
	fmt.Println(generalData)
	fmt.Println(gameData)
}
