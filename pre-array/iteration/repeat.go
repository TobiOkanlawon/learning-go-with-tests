package iteration

import (
	"strings"
)

func OldRepeat(c string, n int) (repeatedString string) {
	for i := 0; i < n; i++ {
		repeatedString += c
	}
	return
}

func Repeat(character string, number_of_times int) (repeated string) {

	return strings.Repeat(character, number_of_times)
}
