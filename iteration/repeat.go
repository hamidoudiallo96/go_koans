package iteration

import (
	"fmt"
	"strings"
)

const repeatCounter = 5

func Repeat(letter string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCounter; i++ {
		repeated.WriteString(letter)
	}

	fmt.Println(repeated)
	return repeated.String()
}
