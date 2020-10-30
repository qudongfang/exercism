package raindrops

import (
	"bytes"
	"strconv"
)

// Convert convert a number into a string that contains raindrop sounds
func Convert(num int) string {
	var b bytes.Buffer

	if 0 == num%3 {
		b.WriteString("Pling")
	}

	if 0 == num%5 {
		b.WriteString("Plang")
	}

	if 0 == num%7 {
		b.WriteString("Plong")
	}

	if 0 == b.Len() {
		b.WriteString(strconv.Itoa(num))
	}

	return b.String()
}
