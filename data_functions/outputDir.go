package dataFunction

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDirectory() string {
	cwd, _ := os.Getwd()
	cwd = filepath.Join(cwd, "Output")

	if _, err := os.Stat(cwd); os.IsNotExist(err) {
		err := os.Mkdir("Output", os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
	return filepath.FromSlash(cwd)
}
