package accumulate

// Accumulate given a collection and an operation to perform on each element of the collection
// returns a new collection containing the result of applying that operation to each element of the input collection
func Accumulate(list []string, fn func(string) string) []string {
	n := make([]string, 0, len(list))

	for _, element := range list {
		n = append(n, fn(element))
	}

	return n
}
