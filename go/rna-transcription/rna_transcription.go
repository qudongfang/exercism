package strand

import "bytes"

// ToRNA DNA to RNA
func ToRNA(dna string) string {
	transMap := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	var b bytes.Buffer
	
	for _, n := range dna {
		if t, ok := transMap[n]; ok {
			b.WriteRune(t)
		} else {
			panic("invalid DNA strand")
		}
	}

	return b.String()
}
