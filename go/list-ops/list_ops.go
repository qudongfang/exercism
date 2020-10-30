package listops

// IntList int list
type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Append given two lists, add all items in the second list to the end of the first list
func (curr IntList) Append(list IntList) IntList {
	curr = append(curr, list...)
	return curr
}

// Concat given a series of lists, combine all items in all lists into one flattened list
func (curr IntList) Concat(list []IntList) IntList {
	for _, l := range list {
		curr = curr.Append(l)
	}

	return curr
}

// Filter given a predicate and a list, return the list of all items for which predicate(item) is True
func (curr IntList) Filter(fn predFunc) IntList {
	result := IntList{}

	for _, item := range curr {
		if fn(item) {
			result = append(result, item)
		}
	}

	return result
}

// Length given a list, return the total number of items within it
func (curr IntList) Length() int {
	return len(curr)
}

// Map given a function and a list, return the list of the results of applying function(item) on all items
func (curr IntList) Map(fn unaryFunc) IntList {
	for idx, item := range curr {
		curr[idx] = fn(item)
	}

	return curr
}

// Foldl given a function, a list, and initial accumulator, fold (reduce) each item
// into the accumulator from the left using function(accumulator, item)
func (curr IntList) Foldl(fn binFunc, accu int) int {
	for _, item := range curr {
		accu = fn(accu, item)
	}

	return accu
}

// Foldr (given a function, a list, and an initial accumulator, fold (reduce) each item
// into the accumulator from the right using function(item, accumulator))
func (curr IntList) Foldr(fn binFunc, accu int) int {
	for i := curr.Length() - 1; i >= 0; i-- {
		accu = fn(curr[i], accu)
	}

	return accu
}

// Reverse given a list, return a list with all the original items, but in reversed order
func (curr IntList) Reverse() IntList {
	n := IntList{}

	for i := curr.Length() - 1; i >= 0; i-- {
		n = append(n, curr[i])
	}

	return n
}
