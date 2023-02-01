package dataFunction

import (
	"regexp"

	"github.com/go-gota/gota/dataframe"
)

func GetData(df dataframe.DataFrame, column string, row int) string {
	if df.Col(column).Elem(row).String() == " " {
		return ""
	} else {
		return df.Col(column).Elem(row).String()
	}
}

func CheckString(input string) bool {
	for _, content := range input {
		if !CheckRune(content) {
			return false
		}
	}
	return true
}

func CheckRune(input rune) bool {
	return regexp.MustCompile(`\d`).MatchString(string(input))
}
