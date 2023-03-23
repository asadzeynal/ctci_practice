package ch1

func oneWay(a string, b string) bool {
	if len(a) == len(b) {
		return checkReplacement(a, b)
	} else if len(a)+1 == len(b) {
		return checkInsertRemove(b, a)
	} else if len(b)+1 == len(a) {
		return checkInsertRemove(a, b)
	}
	return false
}

func checkInsertRemove(longer string, shorter string) bool {
	var i, j, edits int
	for i < len(longer) && j < len(shorter) {
		if longer[i] != shorter[j] {
			edits++
			i++
		}
		i++
		j++
	}
	return edits <= 1
}

func checkReplacement(a string, b string) bool {
	var edits int
	for i := range a {
		if a[i] != b[i] {
			edits++
		}
		if edits > 1 {
			return false
		}
	}
	return true
}
