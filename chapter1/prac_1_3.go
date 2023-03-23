package ch1

func urlify(s string, l int) string {
	arr := []rune(s)
	ls := len(s) - 1
	for i := l - 1; i >= 0; i-- {
		if ls == i {
			break
		}

		if s[i] != ' ' {
			arr[ls] = rune(s[i])
			arr[i] = ' '
			ls--
		} else {
			arr[ls-2] = '%'
			arr[ls-1] = '2'
			arr[ls] = '0'
			ls -= 3
		}
	}
	return string(arr)
}
