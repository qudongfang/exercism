package collatzconjecture

import (
	"errors"
)

// CollatzConjecture Collatz Conjecture or 3x+1 problem
func CollatzConjecture(num int) (int, error) {
	if num < 1 {
		return 0, errors.New("num must >= 1")
	}

	n := 0

	for 1 != num {
		n++

		if 0 == num%2 {
			num /= 2
		} else {
			num = 3*num + 1
		}
	}

	return n, nil
}
