package iteration

import (
	"strings"
)

const repeatCounter = 5

func Repeat(letter string) string {
	var repeated strings.Builder
	for range repeatCounter {
		repeated.WriteString(letter)
	}

	return repeated.String()
}
